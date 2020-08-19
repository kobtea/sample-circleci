package cmd

import (
	"fmt"

	"github.com/kobtea/sample-circleci/pkg/greeting"
	"github.com/spf13/cobra"
)

// greetCmd represents the greet command
var greetCmd = &cobra.Command{
	Use:  "greet [name]",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(greeting.Say(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)
}
