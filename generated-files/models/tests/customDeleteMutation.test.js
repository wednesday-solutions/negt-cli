import get from 'lodash/get';
import { getResponse, mockDBClient, resetAndMockDB } from '@utils/testUtils';
import { modelsTable } from '@utils/testUtils/mockData';

describe('Model graphQL-server-DB mutation tests', () => {
  let dbClient;
  beforeEach(() => {
    dbClient = mockDBClient();
    resetAndMockDB(null, {}, dbClient);
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
