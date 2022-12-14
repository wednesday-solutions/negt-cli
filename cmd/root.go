package cmd

import "github.com/spf13/cobra"

// RootCmd is the command variable of root command negt.
var RootCmd = RootCmdFn()
var version = "1.3.9"

// RootCmd represents the base command when called without any subcommands
func RootCmdFn() *cobra.Command {

	var cmd = &cobra.Command{
		Use:     "negt",
		Version: version,
		Short:   "NodeJS-Express-GraphQL-Template",
		Long: `
	NodeJS-Express-GraphQL-Template. 
	
	It can auto generate graphql models and resolvers by your requirements.. For example:
	If you want to make graphql models and it's resolvers please provide the details for the questions.
	It will auto generate graphql schema and resolvers by itself.
	
	If you are using Node-Express-GraphQL-Template, select "server/gql/models" directory for creating GraphQL Models.
	
	Wednesday Solutions`,
	}
	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	RootCmd.CompletionOptions.DisableDefaultCmd = true

	err := RootCmd.Execute()
	if err != nil {
		return err
	}
	return nil
}
