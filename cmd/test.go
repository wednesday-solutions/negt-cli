package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var testCmd = TestCmd()

func RunTestCmd(cmd *cobra.Command, args []string) error {
	fmt.Print("this command is for testing\n")
	return nil
}

// testCmd represents the test command
func TestCmd() *cobra.Command {

	var testCmd = &cobra.Command{
		Use:   "test",
		Short: "this command is for testing",
		Long:  `this command is for testing`,
		RunE:  RunTestCmd,
	}
	return testCmd
}

func init() {
	gqlgenCmd.AddCommand(testCmd)
}
