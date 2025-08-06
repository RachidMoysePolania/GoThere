/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/RachidMoysePolania/GoThere/server"
	"github.com/spf13/cobra"
)

var relayAddress string
var relayPort string

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Long: `The server command starts the application server.
	This command initializes the server and listens for incoming requests.
	The server is who are going to host the client commands.`,
	Run: func(cmd *cobra.Command, args []string) {
		server.StartServerAgent(relayAddress, relayPort)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	serverCmd.Flags().StringVarP(&relayAddress, "relay-address", "a", "localhost", "Address of the relay server")
	serverCmd.Flags().StringVarP(&relayPort, "relay-port", "p", "443", "Port of the relay server")
	serverCmd.MarkFlagRequired("relay-address")
	serverCmd.MarkFlagRequired("relay-port")
}
