package cmd

import (
	"github.com/ijasMohamad/cobra-cli/gqlgenUtils"

	"github.com/spf13/cobra"
)

// modelCmd represents the model command
var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "model for creating new graphql model",
	Long: `model command for creating new graphql model. For example:

Create gqlmodel by this command, gqlgen model. then provide the 
appropriate answers for the questions.`,
	Run: func(cmd *cobra.Command, args []string) {
		gqlgenUtils.CreateNewModel()
	},
}

func init() {
	gqlgenCmd.AddCommand(modelCmd)
}
