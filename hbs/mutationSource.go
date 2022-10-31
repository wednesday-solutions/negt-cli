package hbs

import "github.com/ijasMohamad/cobra-cli/gqlgenUtils/fileUtils"

func MutationSource(modelName, path, file string, ctx map[string]interface{})  error {
	source := `import { GraphQLID, GraphQLNonNull{{#if graphqlInt}}, {{graphqlInt}}{{/if}}{{#if graphqlString}}, {{graphqlString}}{{/if}} } from 'graphql';
import { GraphQL{{titleSingularModel}} } from './model';
import db from '@database/models';

export const {{singularModel}}MutationFields = {
	id: { type: new GraphQLNonNull(GraphQLID) }{{fieldsWithType fields fieldTypes nullFields}}
};

export const {{singularModel}}Mutation = {
	args: {{singularModel}}MutationFields,
	type: GraphQL{{titleSingularModel}},
	model: db.{{pluralModel}}
};
`
	data, err := GenerateTemplate(source, ctx)
	if err != nil {
		return err
	}
	fileUtils.WriteToFile(path, file, data)

	return nil
}