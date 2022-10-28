import { createConnection } from 'graphql-sequelize';
import db from '@database/models';
import { GraphQLStudent } from './model';
import { sequelizedWhere } from '@server/database/dbUtils';
import { totalConnectionFields } from '@server/utils';

export const studentConnection = createConnection({
	name: 'students',
	target: db.students,
	nodeType: GraphQLStudent,
	before: (findOptions, args, context) => {
		findOptions.include = findOptions.include || [];
		findOptions.where = sequelizedWhere(findOptions.where, args.where);
		return findOptions;
	},
	...totalConnectionFields
});
