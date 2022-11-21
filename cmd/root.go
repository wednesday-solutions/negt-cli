package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "negt",
	Short: "NodeJS-Express-GraphQL-Template",
	Long: `
NodeJS-Express-GraphQL-Template. 

It can auto generate graphql models and resolvers by your requirements.. For example:
If you want to make graphql models and it's resolvers please provide the details for the questions.
It will auto generate graphql schema and resolvers by itself.

Wednesday Solutions`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
NodeJS-Express-GraphQL-Template. 

It can auto generate graphql models and resolvers by your requirements.. For example:
If you want to make graphql models and it's resolvers please provide the details for the questions.
It will auto generate graphql schema and resolvers by itself.

Wednesday Solutions`)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
