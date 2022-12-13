import { GraphQLID, GraphQLNonNull, GraphQLInt, GraphQLString } from 'graphql';
import { GraphQLModel } from './model';
import db from '@database/models';
import { customCreateMutation } from './customCreateMutation';
import { customUpdateMutation } from './customUpdateMutation';
import { customDeleteMutation } from './customDeleteMutation';

export const modelMutationFields = {
  id: { type: new GraphQLNonNull(GraphQLID) },
  name: { type: new GraphQLNonNull(GraphQLString) },
  type: { type: GraphQLInt }
};

export const modelMutation = {
  args: modelMutationFields,
  type: GraphQLModel,
  model: db.models,
  customCreateResolver: customCreateMutation,
  customUpdateResolver: customUpdateMutation,
  customDeleteResolver: customDeleteMutation
};
