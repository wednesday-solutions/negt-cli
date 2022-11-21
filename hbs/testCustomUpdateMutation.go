package hbs

import (
	"fmt"

	"github.com/ijasMohamad/cliApp/gqlgenUtils/fileUtils"
)

func CustomUpdateMutationTestSource(modelName, path, file string, ctx map[string]interface{}) error {

	source := fmt.Sprintf(`import get from 'lodash/get';
import { getResponse, mockDBClient, resetAndMockDB } from '@utils/testUtils';
import { {{pluralModel}}Table } from '@utils/testUtils/mockData';

describe('{{titleSingularModel}} graphQL-server-DB mutation tests', () => {
	let dbClient;
	beforeEach(() => {
		dbClient = mockDBClient();
		resetAndMockDB(null, {}, dbClient);
	});

	const update{{titleSingularModel}}Mutation = %s
mutation {
	update{{titleSingularModel}} (
		id: ${{openingBrace}}{{pluralModel}}Table[0].id}{{{inputStringFieldsWithID fields fieldTypes pluralModel}}}
	) {
		id
	}
}
%s;
	it('should have a mutation to update a new {{singularModel}}', async () => {
		jest.spyOn(dbClient.models.{{pluralModel}}, 'update');
		await getResponse(update{{titleSingularModel}}Mutation).then(response => {
			const result = get(response, 'body.data.update{{titleSingularModel}}');
			expect(result).toBeTruthy();
			expect(dbClient.models.{{pluralModel}}.update.mock.calls.length).toBe(1);
			expect(dbClient.models.{{pluralModel}}.update.mock.calls[0][0]).toEqual({
				id: {{pluralModel}}Table[0].id.toString(){{testFieldsWithID fields pluralModel}}
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
