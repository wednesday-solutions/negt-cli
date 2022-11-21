package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ijasMohamad/cliApp/gqlgenUtils/fileUtils"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create directory gql/models",
	Long:  `This command gqlgen init will create directory for storing gql models`,
	Run: func(cmd *cobra.Command, args []string) {

		path, _ := filepath.Abs(".")
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

		fmt.Println("New directory 'gql/models' created.")
	},
}

func init() {
	gqlgenCmd.AddCommand(initCmd)
}
