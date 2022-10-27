package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/ijasMohamad/cobra-cli/gqlgenUtils"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creating directory gql/models",
	Long:  `This command gqlgen init will create directory for storing gql models`,
	Run: func(cmd *cobra.Command, args []string) {

		path, _ := filepath.Abs(".")
		err := gqlgenUtils.MakeDirectory(path, "gql")
		if err != nil {
			log.Fatal(err)
		}
		err = gqlgenUtils.MakeDirectory(path+"/gql", "models")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("New directory 'gql/models' created.")
	},
}

func init() {
	gqlgenCmd.AddCommand(initCmd)
}
