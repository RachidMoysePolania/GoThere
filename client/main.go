package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// --- CONFIGURATION ---
// Set the address of your relay server on AWS. This MUST match the server's config.
var RELAY_ADDR = ""

// The ID of the server you want to connect to. This MUST match the server's ID.
var TARGET_SERVER_ID = ""

// --- END CONFIGURATION ---

func StartClientAgent(relayAddress, relayPort, targetServerID string) {
	RELAY_ADDR = fmt.Sprintf("%s:%s", relayAddress, relayPort)
	TARGET_SERVER_ID = targetServerID
	log.Println("Starting client...")
	log.Printf("Attempting to connect to relay at %s to reach server '%s'", RELAY_ADDR, TARGET_SERVER_ID)

	// Dial the relay server
	conn, err := net.Dial("tcp", RELAY_ADDR)
	if err != nil {
		log.Fatalf("Could not connect to relay server: %v", err)
	}
	defer conn.Close()

	log.Println("Successfully connected to relay. Requesting connection to server...")

	// --- CONNECTION REQUEST ---
	// Send the command to the relay to connect us to the desired server.
	// The format is "CONNECT <server-id>"
	connectCommand := fmt.Sprintf("CONNECT %s\n", TARGET_SERVER_ID)
	_, err = conn.Write([]byte(connectCommand))
	if err != nil {
		log.Fatalf("Failed to send connect command: %v", err)
	}

	// --- Start a goroutine to read responses from the server ---
	// We do this concurrently so we can be listening for responses
	// at the same time as we are sending messages.
	go readFromServer(conn)

	log.Printf("Connection established with server '%s'. Type a message and press ENTER to send.", TARGET_SERVER_ID)

	// --- Read user input and send it to the server ---
	// This loop will run forever in the main thread.
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Get text from user's console input
		text := scanner.Text()

		// Send the user's text to the server, ensuring it ends with a newline
		_, err := fmt.Fprintf(conn, "%s\n", text)
		if err != nil {
			log.Printf("Failed to send message: %v", err)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading from console: %v", err)
	}
}

// readFromServer continuously reads messages from the connection and prints them.
func readFromServer(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			// If the connection is closed, this error will be caught.
			log.Printf("Connection to server lost: %v", err)
			os.Exit(1) // Exit the client program
		}

		// Print the message received from the server
		fmt.Printf("Server echo: %s", message)
	}
}
