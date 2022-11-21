package hbs

import (
	"fmt"

	"github.com/ijasMohamad/cliApp/gqlgenUtils/fileUtils"
)

func QueryTestSource(modelName, path, file string, ctx map[string]interface{}) error {

	source := fmt.Sprintf(`import get from 'lodash/get';
import { getResponse, mockDBClient, resetAndMockDB } from '@utils/testUtils';
import { {{pluralModel}}Table } from '@server/utils/testUtils/mockData';
	
describe('{{titleSingularModel}} graphQL-server-DB query tests', () => {
	const id = 1;
	const {{singularModel}}One = %s
	query {
		{{singularModel}} (id: ${id}) {
			id{{#each fields}}
			{{this}}{{/each}}
		}
	}
	%s;
	it('should request {{pluralModel}}', async () => {
		const dbClient = mockDBClient();
		resetAndMockDB(null, {}, dbClient);

		jest.spyOn(dbClient.models.{{pluralModel}}, 'findAll').mockImplementation(() => [{{pluralModel}}Table[0]]);

		await getResponse({{singularModel}}One).then(response => {
			const result = get(response, 'body.data.{{singularModel}}');
			expect(result).toBeTruthy();
			expect(result).toEqual(
				expect.objectContaining({
					id: {{pluralModel}}Table[0].id{{testFieldsWithID fields pluralModel}}
				})
			);
			expect(dbClient.models.{{pluralModel}}.findAll.mock.calls.length).toBe(1);
			expect(dbClient.models.{{pluralModel}}.findAll.mock.calls[0][0].include[0].where).toEqual({ id });
			expect(dbClient.models.{{pluralModel}}.findAll.mock.calls[0][0].include[0].model.name).toEqual('{{pluralModel}}');
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