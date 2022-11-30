package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/wednesday-solutions/negt/gqlgenUtils/fileUtils"
)

var initCmd = InitCmd()

func RunInitE(cmd *cobra.Command, args []string) error {
	var (
		SERVER = "server"
		GQL    = "gql"
		MODELS = "models"
	)

	flag := fileUtils.InitUtils()
	path, _ := filepath.Abs(".") // nolint:errcheck

	if flag {
		if !fileUtils.IsExists(path, SERVER) {
			fmt.Println("IsExists")
			_ = fileUtils.MakeDirectory(path, SERVER) // nolint:errcheck
		}
		path = fmt.Sprintf("%s/%s", path, SERVER)
	}

	if !fileUtils.IsExists(path, GQL) {
		_ = fileUtils.MakeDirectory(path, GQL) // nolint:errcheck
	}
	path = fmt.Sprintf("%s/%s", path, GQL)

	if !fileUtils.IsExists(path, MODELS) {
		_ = fileUtils.MakeDirectory(path, MODELS) // nolint:errcheck
		if flag {
			fmt.Printf("New directory '%s/%s/%s' created.", SERVER, GQL, MODELS)
		} else {
			fmt.Printf("New directory '%s/%s' created.", GQL, MODELS)
		}
	} else {
		if flag {
			fmt.Printf("Already initialised '%s/%s/%s'.", SERVER, GQL, MODELS)
		} else {
			fmt.Printf("Already initialised '%s/%s'.", GQL, MODELS)
		}
	}
	return nil
}

// initCmd represents the init command
func InitCmd() *cobra.Command {

	var initCmd = &cobra.Command{
		Use:   "init",
		Short: "Create directory gql/models",
		Long: `This command gqlgen init will create directory for storing gql models
		
	If you are using Node-Express-GraphQL-Template, don't need this command.
	It will create gql directory in the root directory and the models directory inside of gql directory.
	`,
		RunE: RunInitE,
	}
	return initCmd
}

func init() {
	gqlgenCmd.AddCommand(initCmd)
}
