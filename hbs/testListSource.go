package hbs

import (
	"fmt"

	"github.com/ijasMohamad/cliApp/gqlgenUtils/fileUtils"
)

func ListTestSource(modelName, path, file string, ctx map[string]interface{}) error {

	source := fmt.Sprintf(`import get from 'lodash/get';
import { getResponse, mockDBClient, resetAndMockDB } from '@utils/testUtils';
import { {{pluralModel}}Table } from '@server/utils/testUtils/mockData';
	
describe('{{titleSingularModel}} graphQL-server-DB list tests', () => {
	const limit = 1;
	const offset = 0;
	const {{singularModel}}All = %s
	query {
		{{pluralModel}} (limit: ${limit}, offset: ${offset}) {
			edges {
				node {
					id{{#each fields}}
					{{this}}{{/each}}
				}
			}
		}
	}
	%s;
	it('should request {{pluralModel}} with offset and limit', async () => {
		const dbClient = mockDBClient();
		resetAndMockDB(null, {}, dbClient);

		jest.spyOn(dbClient.models.{{pluralModel}}, 'findAll').mockImplementation(() => [{{pluralModel}}Table[0]]);

		await getResponse({{singularModel}}All).then(response => {
      expect(dbClient.models.{{pluralModel}}.findAll.mock.calls.length).toBe(1);
			const result = get(response, 'body.data.{{pluralModel}}.edges[0].node');
			expect(result).toBeTruthy();
			expect(result).toEqual(
				expect.objectContaining({
					id: {{pluralModel}}Table[0].id{{testFieldsWithID fields pluralModel}}
				})
			);
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