package cmd

import (
	"fmt"
	"strings"

	"github.com/Delavalom/qloq/pkg/openai"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qloq [string to print]",
	Short: "Qloq anything to the screen",
	Long: `Qloq is for printing anything back to the screen.
For many years people have printed back to the screen.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(openai.Generate(strings.Join(args, " ")))
	},
}

func Execute() {
	rootCmd.Execute()
}
