import range from 'lodash/range';
import faker from 'faker';

export const studentsTable = range(1, 10).map((_, index) => ({
	id: (index + 1).toString(),
	name: faker.name.firstName(),
	class: 10
}));
