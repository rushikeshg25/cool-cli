package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cool-cli",
	Short: "Command line interface for coolDB",
	Run: ,
}

func Execute(){
	if err:=rootCmd.Execute(); err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
}