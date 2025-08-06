package relay

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
)

var servers = make(map[string]net.Conn)
var lock = &sync.Mutex{}

func StartRelayServer(port string) {
	log.Println("Starting Relay Server...")

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// The first message from a connection determines if it's a server or a client.
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		conn.Close()
		return
	}

	// Example protocol: "REGISTER server-id" or "CONNECT server-id"
	command := string(buffer[:n])
	parts := strings.Split(command, " ")

	if len(parts) == 2 && parts[0] == "REGISTER" {
		serverID := parts[1]
		log.Printf("Server '%s' registered.", serverID)
		lock.Lock()
		// If another server with the same ID is connected, disconnect it.
		if oldConn, ok := servers[serverID]; ok {
			oldConn.Close()
		}
		servers[serverID] = conn
		lock.Unlock()
		// Keep the connection open but don't close it here. The server is now waiting.
	} else if len(parts) == 2 && parts[0] == "CONNECT" {
		serverID := parts[1]
		log.Printf("Client wants to connect to '%s'.", serverID)
		lock.Lock()
		targetConn, ok := servers[serverID]
		lock.Unlock()

		if !ok {
			log.Printf("Server '%s' not found.", serverID)
			conn.Write([]byte("ERROR: Server not found"))
			conn.Close()
			return
		}

		// Stitch the connections together
		log.Printf("Stitching client to server '%s'", serverID)
		go io.Copy(targetConn, conn)
		go io.Copy(conn, targetConn)
	} else {
		log.Println("Invalid command received.")
		conn.Close()
	}
}
