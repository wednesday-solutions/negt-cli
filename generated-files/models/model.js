import { getNode } from '@server/gql/node';
import { GraphQLID, GraphQLNonNull, GraphQLObjectType, GraphQLInt, GraphQLString } from 'graphql';
import { getQueryFields, TYPE_ATTRIBUTES } from '@server/utils/gqlFieldUtils';
import { timestamps } from '@gqlFields/timestamps';

const { nodeInterface } = getNode();

export const modelFields = {
  id: { type: new GraphQLNonNull(GraphQLID) },
  name: { type: new GraphQLNonNull(GraphQLString) },
  type: { type: GraphQLInt }
};

export const GraphQLModel = new GraphQLObjectType({
  name: 'Model',
  interfaces: [nodeInterface],
  fields: () => ({
    ...getQueryFields(modelFields, TYPE_ATTRIBUTES.isNonNull),
    ...timestamps
  })
});
