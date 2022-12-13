import { GraphQLModel } from './model';
import { modelList } from './list';
import { modelQuery } from './query';
import { modelMutation } from './mutation';

const Model = GraphQLModel;
const modelLists = modelList;
const modelQueries = modelQuery;
const modelMutations = modelMutation;

export { Model, modelLists, modelQueries, modelMutations };
