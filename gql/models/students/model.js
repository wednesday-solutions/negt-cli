import { getNode }  from '@server/gql/node';
import { GraphQLID, GraphQLNonNull, GraphQLObjectType, GraphQLInt, GraphQLString } from 'graphql';
import { getQueryFields, TYPE_ATTRIBUTES } from '@server/utils/gqlFieldUtils';

const { nodeInterface } = getNode();

export const studentFields = {
	id: { type: new GraphQLNonNull(GraphQLID) },
	name: { type: GraphQLString },
	class: { type: GraphQLInt }
};

export const GraphQLStudent = new GraphQLObjectType({
	name: 'Student',
	interfaces: [nodeInterface],
	fields: () => ({
		...getQueryFields(studentFields, TYPE_ATTRIBUTES.isNonNull)
	})
});
