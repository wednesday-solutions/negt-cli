import get from 'lodash/get';
import { getResponse, mockDBClient, resetAndMockDB } from '@utils/testUtils';
import { modelsTable } from '@server/utils/testUtils/mockData';

describe('Model graphQL-server-DB list tests', () => {
  const limit = 1;
  const offset = 0;
  const modelAll = `
	query {
		models (limit: ${limit}, offset: ${offset}) {
			edges {
				node {
					id
					name
					type
				}
			}
		}
	}
	`;
  it('should request models with offset and limit', async () => {
    const dbClient = mockDBClient();
    resetAndMockDB(null, {}, dbClient);

    jest.spyOn(dbClient.models.models, 'findAll').mockImplementation(() => [modelsTable[0]]);

    await getResponse(modelAll).then(response => {
      expect(dbClient.models.models.findAll.mock.calls.length).toBe(1);
      const result = get(response, 'body.data.models.edges[0].node');
      expect(result).toBeTruthy();
      expect(result).toEqual(
        expect.objectContaining({
          id: modelsTable[0].id,
          name: modelsTable[0].name,
          type: modelsTable[0].type
        })
      );
    });
  });
});
