package cmd

import (
	"github.com/ijasMohamad/cliApp/gqlgenUtils/modelUtils"
	"github.com/spf13/cobra"
)

// modelCmd represents the model command
var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "Create new graphql model",
	Long: `model command for creating new graphql model. For example:

Create gqlmodel by this command, gqlgen model. then provide the 
appropriate answers for the questions.`,
	Run: func(cmd *cobra.Command, args []string) {
		modelUtils.CreateNewModel()
	},
}

func init() {
	gqlgenCmd.AddCommand(modelCmd)
}
