package hbs

import (
	"fmt"

	"github.com/ijasMohamad/cliApp/gqlgenUtils/fileUtils"
)


func PaginationTestSource(modelName, path, file string, ctx map[string]interface{}) error {

	source := fmt.Sprintf(`import get from 'lodash/get';
import { getResponse, resetAndMockDB } from '@server/utils/testUtils';
import { {{pluralModel}}Table } from '@server/utils/testUtils/mockData';

describe('{{titleSingularModel}} graphql-server-DB pagination tests', () => {
	const {{pluralModel}}Query = %s
	query {
		{{pluralModel}} (first: 1, limit: 1, offset: 0) {
			edges {
				node {
					id
					name
				}
			}
			pageInfo {
				hasNextPage
				hasPreviousPage
				startCursor
				endCursor
			}
			total
		}
	}
	%s;
	it('should have a query to get the {{pluralModel}}', async () => {
		resetAndMockDB(null, {});
		await getResponse({{pluralModel}}Query).then(response => {
			const result = get(response, 'body.data.{{pluralModel}}.edges[0].node');
			expect(result).toEqual(
				expect.objectContaining({
					id: {{pluralModel}}Table[0].id{{testFieldsWithID fields pluralModel}}
				})
			);
		});
	});
	it('should have the correct page info', async () => {
		await getResponse({{pluralModel}}Query).then(response => {
			const result = get(response, 'body.data.{{pluralModel}}.pageInfo');
			expect(result).toEqual(
				expect.objectContaining({
					hasNextPage: true,
					hasPreviousPage: false
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