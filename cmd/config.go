package cmd

import (
	"fmt"
	"log"

	configFile "github.com/Delavalom/qloq/pkg/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use: "config",
	Run: func(cmd *cobra.Command, args []string) {
		msg, err := configFile.Create(args[0])

		if err != nil {
			log.Fatalln(err.Error())
		}

		fmt.Println(msg)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
