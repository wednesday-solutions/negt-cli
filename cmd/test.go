package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wednesday-solutions/negt/hbs"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "this command is for testing",
	Long:  `this command is for testing`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("this command is for testing\n")
		hbs.TestingFunction()
	},
}

func init() {
	gqlgenCmd.AddCommand(testCmd)
}
