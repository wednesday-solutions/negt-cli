package hbs

import (
	"fmt"

	"github.com/ijasMohamad/cliApp/gqlgenUtils/fileUtils"
)

func CustomCreateMutationTestSource(modelName, path, file string, ctx map[string]interface{}) error {

	source := fmt.Sprintf(`import get from 'lodash/get';
import { getResponse, mockDBClient, resetAndMockDB } from '@utils/testUtils';
import { {{pluralModel}}Table } from '@utils/testUtils/mockData';
	
describe('{{titleSingularModel}} graphQL-server-DB mutation tests', () => {
	let dbClient;
	beforeEach(() => {
		dbClient = mockDBClient();
		resetAndMockDB(null, {}, dbClient);
	});
	const create{{titleSingularModel}}Mutation = %s
	mutation {
		create{{titleSingularModel}} ({{{inputStringFieldsWithoutID fields fieldTypes pluralModel}}}
		) {
			id{{#each fields}}
			{{this}}{{/each}}
		}
	}
	%s;
	it('should have a mutation to create a new {{singularModel}}', async () => {
		jest.spyOn(dbClient.models.{{pluralModel}}, 'create');
		await getResponse(create{{titleSingularModel}}Mutation).then(response => {
			const result = get(response, 'body.data.create{{titleSingularModel}}');
			expect(result).toMatchObject({
				id: {{pluralModel}}Table[0].id{{testFieldsWithID fields pluralModel}}
			});
		});
	});
});
`, "`", "`")

	tpl, err := GenerateTemplate(source, ctx)
	if err != nil {
		return err
	}
	fileUtils.WriteToFile(path, file, tpl)

	return nil
}
