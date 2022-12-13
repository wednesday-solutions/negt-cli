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
});
