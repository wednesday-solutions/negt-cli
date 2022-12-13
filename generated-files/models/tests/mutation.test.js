import get from 'lodash/get';
import { getResponse, mockDBClient, resetAndMockDB } from '@utils/testUtils';
import { modelsTable } from '@utils/testUtils/mockData';

describe('Model graphQL-server-DB mutation tests', () => {
  let dbClient;
  beforeEach(() => {
    dbClient = mockDBClient();
    resetAndMockDB(null, {}, dbClient);
  });
  const createModelMutation = `
	mutation {
		createModel (
			name: "${modelsTable[0].name}",
			type: ${modelsTable[0].type}
		) {
			id
			name
			type
		}
	}
	`;
  it('should have a mutation to create a new model', async () => {
    jest.spyOn(dbClient.models.models, 'create');
    await getResponse(createModelMutation).then(response => {
      const result = get(response, 'body.data.createModel');
      expect(result).toMatchObject({
        id: modelsTable[0].id,
        name: modelsTable[0].name,
        type: modelsTable[0].type
      });
    });
  });

  const updateModelMutation = `
	mutation {
		updateModel (
			id: ${modelsTable[0].id},
			name: "${modelsTable[0].name}",
			type: ${modelsTable[0].type}
		) {
			id
		}
	}
	`;
  it('should have a mutation to update a new model', async () => {
    jest.spyOn(dbClient.models.models, 'update');
    await getResponse(updateModelMutation).then(response => {
      const result = get(response, 'body.data.updateModel');
      expect(result).toBeTruthy();
      expect(dbClient.models.models.update.mock.calls.length).toBe(1);
      expect(dbClient.models.models.update.mock.calls[0][0]).toEqual({
        id: modelsTable[0].id.toString(),
        name: modelsTable[0].name,
        type: modelsTable[0].type
      });
    });
  });

  const deleteModelMutation = `
	mutation {
		deleteModel (
			id: ${modelsTable[0].id}
		) {
			id
		}
	}
	`;
  it('should have a mutation to delete a model', async () => {
    jest.spyOn(dbClient.models.models, 'destroy');
    await getResponse(deleteModelMutation).then(response => {
      const result = get(response, 'body.data.deleteModel');
      expect(result).toBeTruthy();
      expect(dbClient.models.models.destroy.mock.calls.length).toBe(1);
      expect(dbClient.models.models.destroy.mock.calls[0][0]).toEqual({
        where: {
          deletedAt: null,
          id: parseInt(modelsTable[0].id)
        }
      });
    });
  });
});
