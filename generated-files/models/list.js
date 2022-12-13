import { createConnection } from 'graphql-sequelize';
import db from '@database/models';
import { GraphQLModel } from './model';
import { sequelizedWhere } from '@server/database/dbUtils';
import { totalConnectionFields } from '@server/utils';

export const modelConnection = createConnection({
  name: 'models',
  target: db.models,
  nodeType: GraphQLModel,
  before: (findOptions, args, context) => {
    findOptions.include = findOptions.include || [];
    findOptions.where = sequelizedWhere(findOptions.where, args.where);
    return findOptions;
  },
  ...totalConnectionFields
});

export const modelList = {
  list: {
    ...modelConnection,
    resolve: modelConnection.resolve,
    type: modelConnection.connectionType,
    args: modelConnection.connectionArgs
  },
  model: db.models
};
