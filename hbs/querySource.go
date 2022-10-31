package hbs

import "github.com/ijasMohamad/cobra-cli/gqlgenUtils/fileUtils"

func QuerySource(modelName, path, file string, ctx map[string]interface{})  error {

	source := `import { GraphQLInt, GraphQLNonNull } from 'graphql';
import { GraphQL{{titleSingularModel}} } from './model';
import { {{singularModel}}Connection } from './list';
import db from '@database/models';

export const {{titleSingularModel}}Queries = {
	args: {
		id: {
			type: new GraphQLNonNull(GraphQLInt)
		}
	},
	query: {
		type: GraphQL{{titleSingularModel}}
	},
	list: {
		...{{singularModel}}Connection,
		resolve: {{singularModel}}Connection.resolve,
		type: {{singularModel}}Connection.connectionType,
		args: {{singularModel}}Connection.connectionArgs
	},
	model: db.{{pluralModel}}
};
`
	tpl, err := GenerateTemplate(source, ctx)
	if err != nil {
		return err
	}
	fileUtils.WriteToFile(path, file, tpl)

	return nil
}