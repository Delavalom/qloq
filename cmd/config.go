package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use: "config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args[0])
		err := os.Setenv("API_KEY", args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("API_KEY setup successfully")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
