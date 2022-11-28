import get from 'lodash/get';
import { getResponse, mockDBClient, resetAndMockDB } from '@utils/testUtils';
import { studentsTable } from '@utils/testUtils/mockData';
	
describe('Student graphQL-server-DB mutation tests', () => {
  let dbClient;
  beforeEach(() => {
    dbClient = mockDBClient();
    resetAndMockDB(null, {}, dbClient);
  });
	const deleteStudentMutation = `
	mutation {
		deleteStudent (
			id: ${studentsTable[0].id}
		) {
			id
		}
	}
	`;
	it('should have a mutation to delete a student', async () => {
		jest.spyOn(dbClient.models.students, 'destroy');
		await getResponse(deleteStudentMutation).then(response => {
			const result = get(response, 'body.data.deleteStudent');
			expect(result).toBeTruthy();
			expect(dbClient.models.students.destroy.mock.calls.length).toBe(1);
			expect(dbClient.models.students.destroy.mock.calls[0][0]).toEqual({
				where: {
					deletedAt: null,
					id: parseInt(studentsTable[0].id)
				}
			});
		});
	});
});
