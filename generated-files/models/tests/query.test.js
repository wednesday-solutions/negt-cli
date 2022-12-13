import get from 'lodash/get';
import { getResponse, mockDBClient, resetAndMockDB } from '@utils/testUtils';
import { modelsTable } from '@server/utils/testUtils/mockData';

describe('Model graphQL-server-DB query tests', () => {
  const id = 1;
  const modelOne = `
	query {
		model (id: ${id}) {
			id
			name
			type
		}
	}
	`;
  it('should request models', async () => {
    const dbClient = mockDBClient();
    resetAndMockDB(null, {}, dbClient);

    jest.spyOn(dbClient.models.models, 'findAll').mockImplementation(() => [modelsTable[0]]);

    await getResponse(modelOne).then(response => {
      const result = get(response, 'body.data.model');
      expect(result).toBeTruthy();
      expect(result).toEqual(
        expect.objectContaining({
          id: modelsTable[0].id,
          name: modelsTable[0].name,
          type: modelsTable[0].type
        })
      );
      expect(dbClient.models.models.findAll.mock.calls.length).toBe(1);
      expect(dbClient.models.models.findAll.mock.calls[0][0].include[0].where).toEqual({ id });
      expect(dbClient.models.models.findAll.mock.calls[0][0].include[0].model.name).toEqual('models');
    });
  });
});
