package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ijasMohamad/negt/gqlgenUtils/fileUtils"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create directory gql/models",
	Long:  `This command gqlgen init will create directory for storing gql models
	
If you are using Node-Express-GraphQL-Template, don't need this command.
It will create gql directory in the root directory and the models directory inside of gql directory.
	`,
	Run: func(cmd *cobra.Command, args []string) {

		flag := fileUtils.InitUtils()
		var path string
		if flag {

			path, _ = filepath.Abs("..")
			if !fileUtils.IsExists(path, "server") {
				_ = fileUtils.MakeDirectory(path, "server")
			}
			path = fmt.Sprintf("%s/%s", path, "server")

		} else {
			path, _ = filepath.Abs(".")
		}
		err := fileUtils.MakeDirectory(path, "gql")
		if err != nil {
			if os.IsExist(err) {

				err = fileUtils.MakeDirectory(path+"/gql", "models")
				if err != nil {
					if os.IsExist(err) {
						fmt.Println("Already initialized.")
					}
					fmt.Println("try 'negt gqlgen help'.")
					os.Exit(1)
				}

				fmt.Println("New directory 'gql/models' created.")
				return
			}
			fmt.Println("try 'negt gqlgen help'.")
			os.Exit(1)
		}

		err = fileUtils.MakeDirectory(path+"/gql", "models")
		if err != nil {
			if os.IsExist(err) {
				fmt.Println("Already initialized.")
			}
			fmt.Println("try 'negt gqlgen help'.")
			os.Exit(1)
		}

		if flag {
			fmt.Println("New directory 'server/gql/models' created.")
			
		} else {
			fmt.Println("New directory 'gql/models' created.")
		}
	},
}

func init() {
	gqlgenCmd.AddCommand(initCmd)
}
