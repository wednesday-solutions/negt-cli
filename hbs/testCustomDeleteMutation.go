package hbs

import (
	"fmt"

	"github.com/ijasMohamad/cliApp/gqlgenUtils/fileUtils"
)

func CustomDeleteMutationTestSource(modelName, path, file string, ctx map[string]interface{}) error {

	source := fmt.Sprintf(`import get from 'lodash/get';
import { getResponse, mockDBClient, resetAndMockDB } from '@utils/testUtils';
import { {{pluralModel}}Table } from '@utils/testUtils/mockData';
	
describe('{{titleSingularModel}} graphQL-server-DB mutation tests', () => {
  let dbClient;
  beforeEach(() => {
    dbClient = mockDBClient();
    resetAndMockDB(null, {}, dbClient);
  });
	const delete{{titleSingularModel}}Mutation = %s
	mutation {
		delete{{titleSingularModel}} (
			id: ${{openingBrace}}{{pluralModel}}Table[0].id}
		) {
			id
		}
	}
	%s;
	it('should have a mutation to delete a {{singularModel}}', async () => {
		jest.spyOn(dbClient.models.{{pluralModel}}, 'destroy');
		await getResponse(delete{{titleSingularModel}}Mutation).then(response => {
			const result = get(response, 'body.data.delete{{titleSingularModel}}');
			expect(result).toBeTruthy();
			expect(dbClient.models.{{pluralModel}}.destroy.mock.calls.length).toBe(1);
			expect(dbClient.models.{{pluralModel}}.destroy.mock.calls[0][0]).toEqual({
				where: {
					deletedAt: null,
					id: parseInt({{pluralModel}}Table[0].id)
				}
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
