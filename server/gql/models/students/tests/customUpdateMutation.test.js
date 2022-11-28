import get from 'lodash/get';
import { getResponse, mockDBClient, resetAndMockDB } from '@utils/testUtils';
import { studentsTable } from '@utils/testUtils/mockData';

describe('Student graphQL-server-DB mutation tests', () => {
	let dbClient;
	beforeEach(() => {
		dbClient = mockDBClient();
		resetAndMockDB(null, {}, dbClient);
	});

	const updateStudentMutation = `
mutation {
	updateStudent (
		id: ${studentsTable[0].id},
			name: "${studentsTable[0].name}",
			class: ${studentsTable[0].class}
	) {
		id
	}
}
`;
	it('should have a mutation to update a new student', async () => {
		jest.spyOn(dbClient.models.students, 'update');
		await getResponse(updateStudentMutation).then(response => {
			const result = get(response, 'body.data.updateStudent');
			expect(result).toBeTruthy();
			expect(dbClient.models.students.update.mock.calls.length).toBe(1);
			expect(dbClient.models.students.update.mock.calls[0][0]).toEqual({
				id: studentsTable[0].id.toString(),
				name: studentsTable[0].name,
				class: studentsTable[0].class
			});
		});
	});
});
