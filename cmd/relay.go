/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/RachidMoysePolania/GoThere/relay"
	"github.com/spf13/cobra"
)

var port string

// relayCmd represents the relay command
var relayCmd = &cobra.Command{
	Use:   "relay",
	Short: "Start the relay server",
	Long:  `The relay command starts the relay server which listens for incoming requests and processes them accordingly.`,
	Run: func(cmd *cobra.Command, args []string) {
		relay.StartRelayServer(port)
	},
}

func init() {
	rootCmd.AddCommand(relayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// relayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// relayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	relayCmd.Flags().StringVarP(&port, "port", "p", "443", "Port to listen on for relay server")
	relayCmd.MarkFlagRequired("port")
	relayCmd.Flags().Lookup("port").NoOptDefVal = "443"
}
