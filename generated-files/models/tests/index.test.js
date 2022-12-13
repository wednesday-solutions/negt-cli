import { GraphQLModel } from './model';
import { modelList } from './list';
import { modelQuery } from './query';
import { modelMutation } from './mutation';

describe('Checking importing is succesfull', () => {
  it('Should check imports', async () => {
    expect(GraphQLModel).toBeTruthy();
    expect(Object.keys(GraphQLModel).length > 0).toBe(true);
    expect(modelList).toBeTruthy();
    expect(Object.keys(modelList).length > 0).toBe(true);
    expect(modelQuery).toBeTruthy();
    expect(Object.keys(modelQuery).length > 0).toBe(true);
    expect(modelMutation).toBeTruthy();
    expect(Object.keys(modelMutation).length > 0).toBe(true);
  });
});
