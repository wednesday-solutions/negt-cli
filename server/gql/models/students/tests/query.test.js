import get from 'lodash/get';
import { getResponse, mockDBClient, resetAndMockDB } from '@utils/testUtils';
import { studentsTable } from '@server/utils/testUtils/mockData';
	
describe('Student graphQL-server-DB query tests', () => {
	const id = 1;
	const studentOne = `
	query {
		student (id: ${id}) {
			id
			name
			class
		}
	}
	`;
	it('should request students', async () => {
		const dbClient = mockDBClient();
		resetAndMockDB(null, {}, dbClient);

		jest.spyOn(dbClient.models.students, 'findAll').mockImplementation(() => [studentsTable[0]]);

		await getResponse(studentOne).then(response => {
			const result = get(response, 'body.data.student');
			expect(result).toBeTruthy();
			expect(result).toEqual(
				expect.objectContaining({
					id: studentsTable[0].id,
				name: studentsTable[0].name,
				class: studentsTable[0].class
				})
			);
			expect(dbClient.models.students.findAll.mock.calls.length).toBe(1);
			expect(dbClient.models.students.findAll.mock.calls[0][0].include[0].where).toEqual({ id });
			expect(dbClient.models.students.findAll.mock.calls[0][0].include[0].model.name).toEqual('students');
		});
	});
});
