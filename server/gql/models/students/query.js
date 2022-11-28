import { GraphQLInt, GraphQLNonNull } from 'graphql';
import { GraphQLStudent } from './model';
import db from '@database/models';

export const studentQuery = {
	args: {
		id: {
			type: new GraphQLNonNull(GraphQLInt)
		}
	},
	query: {
		type: GraphQLStudent
	},
	model: db.students
};
