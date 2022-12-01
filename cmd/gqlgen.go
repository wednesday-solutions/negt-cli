package cmd

import (
	"github.com/spf13/cobra"
)

var gqlgenCmd = GqlgenCmd()

// alias for 'gqlgen model'
// var subModelCmd = SubModelCmd()

// GqlgenCmd represents the gqlgen command
func GqlgenCmd() *cobra.Command {

	var gqlgenCmd = &cobra.Command{
		Use:   "gqlgen",
		Short: "Generate graphql models",
		Long: `It is for generating graphql models.
		
	If you want to make graphql models and it's resolvers please provide the details for the questions.
	It will auto generate graphql schema and resolvers by itself.
	
	If you are using Node-Express-GraphQL-Template, select "server/gql/models" directory for creating GraphQL Models.
	`,
	}
	return gqlgenCmd
}

func init() {

	rootCmd.AddCommand(gqlgenCmd)

	// negt gqlgen only give suggestions
	gqlgenCmd.Flags().BoolP("help", "h", false, "Help for gqlgen")

	// alias for 'gqlgen model'
	// rootCmd.AddCommand(subModelCmd)
}
