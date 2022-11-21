package hbs

import "github.com/ijasMohamad/cliApp/gqlgenUtils/fileUtils"

func MutationSource(modelName, path, file string, ctx map[string]interface{})  error {
	source := `import { GraphQLID, GraphQLNonNull{{#if graphqlInt}}, {{graphqlInt}}{{/if}}{{#if graphqlString}}, {{graphqlString}}{{/if}}{{#if graphqlFloat}}, {{graphqlFloat}}{{/if}}{{#if graphqlBoolean}}, {{graphqlBoolean}}{{/if}}{{#if graphqlDateTime}}, {{graphqlDateTime}}{{/if}} } from 'graphql';
import { GraphQL{{titleSingularModel}} } from './model';
import db from '@database/models';{{{customMutationImports customMutation}}} 

export const {{singularModel}}MutationFields = {
	id: { type: new GraphQLNonNull(GraphQLID) }{{fieldsWithType fields fieldTypes nullFields}}
};

export const {{singularModel}}Mutation = {
	args: {{singularModel}}MutationFields,
	type: GraphQL{{titleSingularModel}},
	model: db.{{pluralModel}}{{customMutations customMutation}}
};
`
	tpl, err := GenerateTemplate(source, ctx)
	if err != nil {
		return err
	}
	fileUtils.WriteToFile(path, file, tpl)

	return nil
}