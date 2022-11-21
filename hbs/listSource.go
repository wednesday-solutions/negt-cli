package hbs

import "github.com/ijasMohamad/cliApp/gqlgenUtils/fileUtils"

func ListSource(modelName, path, file string, ctx map[string]interface{})  error {
	source := `import { createConnection } from 'graphql-sequelize';
import db from '@database/models';
import { GraphQL{{titleSingularModel}} } from './model';
import { sequelizedWhere } from '@server/database/dbUtils';
import { totalConnectionFields } from '@server/utils';

export const {{singularModel}}Connection = createConnection({
	name: '{{pluralModel}}',
	target: db.{{pluralModel}},
	nodeType: GraphQL{{titleSingularModel}},
	before: (findOptions, args, context) => {
		findOptions.include = findOptions.include || [];
		findOptions.where = sequelizedWhere(findOptions.where, args.where);
		return findOptions;
	},
	...totalConnectionFields
});

export const {{singularModel}}List = {
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