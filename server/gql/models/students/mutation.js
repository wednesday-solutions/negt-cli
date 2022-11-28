import { GraphQLID, GraphQLNonNull, GraphQLInt, GraphQLString } from 'graphql';
import { GraphQLStudent } from './model';
import db from '@database/models';
import { customCreateMutation } from './customCreateMutation';  
import { customUpdateMutation } from './customUpdateMutation';  
import { customDeleteMutation } from './customDeleteMutation'; 

export const studentMutationFields = {
	id: { type: new GraphQLNonNull(GraphQLID) },
	name: { type: new GraphQLNonNull(GraphQLString) },
	class: { type: new GraphQLNonNull(GraphQLInt) }
};

export const studentMutation = {
	args: studentMutationFields,
	type: GraphQLStudent,
	model: db.students,
	customCreateResolver: customCreateMutation,
	customUpdateResolver: customUpdateMutation,
	customDeleteResolver: customDeleteMutation
};
