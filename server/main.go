package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/google/uuid"
)

// --- CONFIGURATION ---
// Set the address of your relay server on AWS.
var RELAY_ADDR = ""

// Unique ID for this server. In a real app, this might be from a config file.
var SERVER_ID = uuid.New().String()

// --- END CONFIGURATION ---

func StartServerAgent(relayAddress, port string) {
	RELAY_ADDR = fmt.Sprintf("%s:%s", relayAddress, port)
	log.Printf("Starting server agent with ID: %s", SERVER_ID)
	log.Printf("Attempting to connect to relay server at: %s", RELAY_ADDR)

	// This infinite loop ensures the server tries to reconnect if it gets disconnected.
	for {
		connectAndRegister()
		log.Println("Disconnected from relay. Reconnecting in 5 seconds...")
		time.Sleep(5 * time.Second)
	}
}

// connectAndRegister dials the relay, registers the server, and handles the connection.
func connectAndRegister() {
	// Dial the relay server
	conn, err := net.Dial("tcp", RELAY_ADDR)
	if err != nil {
		log.Printf("Error connecting to relay: %v", err)
		return // Will trigger the reconnect loop in main()
	}
	defer conn.Close()

	log.Println("Successfully connected to relay server.")

	// --- REGISTRATION ---
	// The first thing we do is send our registration command.
	// The format is "REGISTER <server-id>"
	registerCommand := fmt.Sprintf("REGISTER %s\n", SERVER_ID)
	_, err = conn.Write([]byte(registerCommand))
	if err != nil {
		log.Printf("Failed to send registration command: %v", err)
		return
	}

	log.Printf("Server registered with ID '%s'. Waiting for client commands...", SERVER_ID)

	// --- COMMAND HANDLING ---
	// This function will now handle all future communication on this connection.
	// It will only return when the connection is closed or an error occurs.
	handleCommands(conn)
}

// handleCommands is the main loop for reading data from the client (via the relay).
func handleCommands(conn net.Conn) {
	// Use a buffered reader for efficiency
	reader := bufio.NewReader(conn)

	for {
		// In the full application, you would decode JSON commands here.
		// For this example, we'll just read and print any incoming message.
		message, err := reader.ReadString('\n')
		if err != nil {
			// This error typically means the client disconnected or the connection broke.
			log.Printf("Connection closed: %v", err)
			return // Exit the function to trigger a reconnect.
		}

		// TODO: This is where you will add your logic from the previous answer.
		// You'll decode the JSON command and use a switch statement to handle:
		// - Filesystem access
		// - Remote control events, etc.
		log.Printf("Received command: %s", message)

		// Example of sending a response back
		// In a real app, this response would be a JSON object.
		conn.Write([]byte("Command received and processed.\n"))
	}
}
