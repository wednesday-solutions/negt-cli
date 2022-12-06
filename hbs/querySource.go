package hbs

import "github.com/wednesday-solutions/negt/gqlgenUtils/fileUtils"

// QuerySource is the source function for query file.
func QuerySource(modelName, path, file string, ctx map[string]interface{}) error {

	source := `import { GraphQLInt, GraphQLNonNull } from 'graphql';
import { GraphQL{{titleSingularModel}} } from './model';
import db from '@database/models';

export const {{singularModel}}Query = {
	args: {
		id: {
			type: new GraphQLNonNull(GraphQLInt)
		}
	},
	query: {
		type: GraphQL{{titleSingularModel}}
	},
	model: db.{{pluralModel}}
};
`
	tpl, err := GenerateTemplate(source, ctx)
	if err != nil {
		return err
	}
	fileUtils.WriteToFile(path, file, tpl) // nolint:errcheck

	return nil
}
