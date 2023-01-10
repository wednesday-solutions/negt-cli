package cmd

import (
	"github.com/spf13/cobra"
)

// GqlgenCmd is the command variable of gqlgen.
var GqlgenCmd = GqlgenCmdFn()

// GqlgenCmd represents the gqlgen command
func GqlgenCmdFn() *cobra.Command {

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

	RootCmd.AddCommand(GqlgenCmd)

	// negt gqlgen only give suggestions
	GqlgenCmd.Flags().BoolP("help", "h", false, "Help for gqlgen")
}
