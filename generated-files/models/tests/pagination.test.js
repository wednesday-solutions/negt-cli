import get from 'lodash/get';
import { getResponse, resetAndMockDB } from '@server/utils/testUtils';
import { modelsTable } from '@server/utils/testUtils/mockData';

describe('Model graphql-server-DB pagination tests', () => {
  const modelsQuery = `
	query {
		models (first: 1, limit: 1, offset: 0) {
			edges {
				node {
					id
					name
				}
			}
			pageInfo {
				hasNextPage
				hasPreviousPage
				startCursor
				endCursor
			}
			total
		}
	}
	`;
  it('should have a query to get the models', async () => {
    resetAndMockDB(null, {});
    await getResponse(modelsQuery).then(response => {
      const result = get(response, 'body.data.models.edges[0].node');
      expect(result).toEqual(
        expect.objectContaining({
          id: modelsTable[0].id,
          name: modelsTable[0].name,
          type: modelsTable[0].type
        })
      );
    });
  });
  it('should have the correct page info', async () => {
    await getResponse(modelsQuery).then(response => {
      const result = get(response, 'body.data.models.pageInfo');
      expect(result).toEqual(
        expect.objectContaining({
          hasNextPage: true,
          hasPreviousPage: false
        })
      );
    });
  });
});
