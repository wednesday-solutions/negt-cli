/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/ijasMohamad/cobra-cli/gqlgenUtils"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creating directory gql/models",
	Long: `This command gqlgen init will create directory for storing gql models`,
	Run: func(cmd *cobra.Command, args []string) {

		path, _ := filepath.Abs(".")
		err := gqlgenUtils.MakeDirectory(path, "gql")
		if err != nil {
			log.Fatal(err)
		}
		err = gqlgenUtils.MakeDirectory(path + "/gql", "models")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("New directory 'gql/models' created.")
	},
}


func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
