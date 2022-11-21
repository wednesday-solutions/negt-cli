package hbs

import (
	"fmt"

	"github.com/ijasMohamad/cliApp/gqlgenUtils/fileUtils"
)


func ModelTestSource(modelName, path, file string, ctx map[string]interface{}) error {

	source := fmt.Sprintf(`import get from 'lodash/get';
import { graphqlSync, GraphQLSchema } from 'graphql';
import { createFieldsWithType, expectSameTypeNameOrKind } from '@utils/testUtils';
import { QueryRoot } from '../../../queries';
import { MutationRoot } from '../../../mutations';
import { {{singularModel}}Fields } from '@gql/models/{{pluralModel}}';

const schema = new GraphQLSchema({ query: QueryRoot, mutation: MutationRoot });

let fields = [];

fields = createFieldsWithType({ ...{{singularModel}}Fields });

const query = %s
	{
		__type(name: "{{titleSingularModel}}") {
			name
			kind
			fields {
				name
				type {
					name
					kind
				}
			}
		}
	}
%s;
describe('{{titleSingularModel}} introspection tests', () => {
	it('should have the correct fields and types', async () => {
		const result = await graphqlSync({ schema, source: query });
		const {{singularModel}}FieldTypes = get(result, 'data.__type.fields');
		const hasCorrectFieldTypes = expectSameTypeNameOrKind({{singularModel}}FieldTypes, fields);
		expect(hasCorrectFieldTypes).toBeTruthy();
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