import { GraphQLInt, GraphQLNonNull } from 'graphql';
import { GraphQLModel } from './model';
import db from '@database/models';

export const modelQuery = {
  args: {
    id: {
      type: new GraphQLNonNull(GraphQLInt)
    }
  },
  query: {
    type: GraphQLModel
  },
  model: db.models
};
