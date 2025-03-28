package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().String("host", "localhost", "hostname or ip address of the DiceDB server")
	rootCmd.PersistentFlags().Int("port", 3040, "port number of the DiceDB server")
}

var rootCmd = &cobra.Command{
	Use:   "cool-cli",
	Short: "Command line interface for coolDB",
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetInt("port")
		fmt.Println("Starting cool-cli")
		fmt.Println("Host: ", host)
		fmt.Println("Port: ", port)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
