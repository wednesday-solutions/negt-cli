/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ijasMohamad/cobra-cli/gqlgenUtils"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "model for creating new graphql model",
	Long: `A longer description model command for creating new graphql model. For example:

Create gqlmodel by this command, gqlgen model.`,
	Run: func(cmd *cobra.Command, args []string) {
		gqlgenUtils.CreateNewModel()
	},
}

func init() {
	rootCmd.AddCommand(modelCmd)
}
