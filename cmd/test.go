package cmd

import (
	"fmt"

	"github.com/ijasMohamad/cliApp/hbs"
	"github.com/spf13/cobra"
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
