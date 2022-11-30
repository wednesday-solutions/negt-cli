package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wednesday-solutions/negt/gqlgenUtils/modelUtils"
)

var modelCmd = ModelCmd()

func RunModelE(cmd *cobra.Command, args []string) error {
	err := modelUtils.CreateNewModel()
	if err != nil {
		return err
	}
	return nil
}

// modelCmd represents the model command
func ModelCmd() *cobra.Command {

	var modelCmd = &cobra.Command{
		Use:     "model",
		Aliases: []string{"gqlmodel"},
		Short:   "Create new graphql model",
		Long: `model command for creating new graphql model. For example:
	
	Create gqlmodel by this command, gqlgen model. then provide the 
	appropriate answers for the questions.`,
		RunE: RunModelE,
	}
	return modelCmd
}

func init() {
	gqlgenCmd.AddCommand(modelCmd)
}
