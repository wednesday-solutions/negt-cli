import get from 'lodash/get';
import { getResponse, mockDBClient, resetAndMockDB } from '@utils/testUtils';
import { studentsTable } from '@server/utils/testUtils/mockData';
	
describe('Student graphQL-server-DB list tests', () => {
	const limit = 1;
	const offset = 0;
	const studentAll = `
	query {
		students (limit: ${limit}, offset: ${offset}) {
			edges {
				node {
					id
					name
					class
				}
			}
		}
	}
	`;
	it('should request students with offset and limit', async () => {
		const dbClient = mockDBClient();
		resetAndMockDB(null, {}, dbClient);

		jest.spyOn(dbClient.models.students, 'findAll').mockImplementation(() => [studentsTable[0]]);

		await getResponse(studentAll).then(response => {
      expect(dbClient.models.students.findAll.mock.calls.length).toBe(1);
			const result = get(response, 'body.data.students.edges[0].node');
			expect(result).toBeTruthy();
			expect(result).toEqual(
				expect.objectContaining({
					id: studentsTable[0].id,
				name: studentsTable[0].name,
				class: studentsTable[0].class
				})
			);
		});
	});
});
