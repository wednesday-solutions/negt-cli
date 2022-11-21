package cmd

import (
	"github.com/spf13/cobra"
)

// gqlgenCmd represents the gqlgen command
var gqlgenCmd = &cobra.Command{
	Use:   "gqlgen",
	Short: "Generate graphql models",
	Long: `It is for generating graphql models.
	
If you want to make graphql models and it's resolvers please provide the details for the questions.
It will auto generate graphql schema and resolvers by itself.`,
}

func init() {
	rootCmd.AddCommand(gqlgenCmd)

	// negt gqlgen only give suggestions
	gqlgenCmd.Flags().BoolP("help", "h", false, "Help for gqlgen")
}
