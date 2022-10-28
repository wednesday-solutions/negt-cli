import { GraphQLID, GraphQLNonNull, GraphQLInt, GraphQLString } from 'graphql';
import { GraphQLStudent } from './model';
import db from '@database/models';

export const studentMutationFields = {
	id: { type: new GraphQLNonNull(GraphQLID) },
	name: { type: GraphQLString },
	class: { type: GraphQLInt }
};

export const studentMutation = {
	args: studentMutationFields,
	type: GraphQLStudent,
	model: db.students
};
