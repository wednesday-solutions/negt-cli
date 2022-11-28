import get from 'lodash/get';
import { getResponse, mockDBClient, resetAndMockDB } from '@utils/testUtils';
import { studentsTable } from '@utils/testUtils/mockData';
	
describe('Student graphQL-server-DB mutation tests', () => {
	let dbClient;
	beforeEach(() => {
		dbClient = mockDBClient();
		resetAndMockDB(null, {}, dbClient);
	});
	const createStudentMutation = `
	mutation {
		createStudent (
			name: "${studentsTable[0].name}",
			class: ${studentsTable[0].class}
		) {
			id
			name
			class
		}
	}
	`;
	it('should have a mutation to create a new student', async () => {
		jest.spyOn(dbClient.models.students, 'create');
		await getResponse(createStudentMutation).then(response => {
			const result = get(response, 'body.data.createStudent');
			expect(result).toMatchObject({
				id: studentsTable[0].id,
				name: studentsTable[0].name,
				class: studentsTable[0].class
			});
		});
	});
});
