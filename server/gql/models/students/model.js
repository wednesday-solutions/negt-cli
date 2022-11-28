import { getNode }  from '@server/gql/node';
import { GraphQLID, GraphQLNonNull, GraphQLObjectType, GraphQLInt, GraphQLString } from 'graphql';
import { getQueryFields, TYPE_ATTRIBUTES } from '@server/utils/gqlFieldUtils';
import { timestamps } from '@gqlFields/timestamps';

const { nodeInterface } = getNode();

export const studentFields = {
	id: { type: new GraphQLNonNull(GraphQLID) },
	name: { type: new GraphQLNonNull(GraphQLString) },
	class: { type: new GraphQLNonNull(GraphQLInt) }
};

export const GraphQLStudent = new GraphQLObjectType({
	name: 'Student',
	interfaces: [nodeInterface],
	fields: () => ({
		...getQueryFields(studentFields, TYPE_ATTRIBUTES.isNonNull),
		...timestamps
	})
});
