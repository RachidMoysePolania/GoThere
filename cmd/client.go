/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/RachidMoysePolania/GoThere/client"
	"github.com/spf13/cobra"
)

var targetServerID string

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Start the client agent",
	Long: `The client command starts the client agent that connects to a relay server.
It requires the relay address and target server ID to establish a connection.`,
	Run: func(cmd *cobra.Command, args []string) {
		client.StartClientAgent(relayAddress, relayPort, targetServerID)
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	clientCmd.Flags().StringVarP(&relayAddress, "relay-address", "a", "localhost", "Address of the relay server")
	clientCmd.Flags().StringVarP(&relayPort, "relay-port", "p", "8080", "Port of the relay server")
	clientCmd.Flags().StringVarP(&targetServerID, "target-server-id", "t", "my-office-pc", "ID of the target server to connect to")
	clientCmd.MarkFlagRequired("relay-address")
	clientCmd.MarkFlagRequired("relay-port")
	clientCmd.MarkFlagRequired("target-server-id")
}
