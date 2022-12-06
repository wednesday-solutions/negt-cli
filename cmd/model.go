package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wednesday-solutions/negt/gqlgenUtils/modelUtils"
)

// ModelCmd is the command variable of model command.
var ModelCmd = ModelCmdFn()

// RunModelE represents the run function for model command
func RunModelE(*cobra.Command, []string) error {
	err := modelUtils.CreateNewModel()
	if err != nil {
		return err
	}
	return nil
}

// ModelCmd represents the model command
func ModelCmdFn() *cobra.Command {

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
	GqlgenCmd.AddCommand(ModelCmd)
}
