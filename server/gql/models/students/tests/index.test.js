import { GraphQLStudent } from './model';
import { studentList } from './list';
import { studentQuery } from './query';
import { studentMutation } from './mutation';
	
descibe('Checking importing is succesfull', () => {
	it('Should check imports', async () => {
		expect(GraphQLStudent).toBeTruthy();
		expect(Object.keys(GraphQLStudent).length > 0).toBe(true);
		expect(studentList).toBeTruthy();
		expect(Object.keys(studentList).length > 0).toBe(true);
		expect(studentQuery).toBeTruthy();
		expect(Object.keys(studentQuery).length > 0).toBe(true);
		expect(studentMutation).toBeTruthy();
		expect(Object.keys(studentMutation).length > 0).toBe(true);
	});
});
