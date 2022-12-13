import get from 'lodash/get';
import { graphqlSync, GraphQLSchema } from 'graphql';
import { createFieldsWithType, expectSameTypeNameOrKind } from '@utils/testUtils';
import { QueryRoot } from '../../../queries';
import { MutationRoot } from '../../../mutations';
import { modelFields } from '@gql/models/models';

const schema = new GraphQLSchema({ query: QueryRoot, mutation: MutationRoot });

let fields = [];

fields = createFieldsWithType({ ...modelFields });

const query = `
	{
		__type(name: "Model") {
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
`;
describe('Model introspection tests', () => {
  it('should have the correct fields and types', async () => {
    const result = await graphqlSync({ schema, source: query });
    const modelFieldTypes = get(result, 'data.__type.fields');
    const hasCorrectFieldTypes = expectSameTypeNameOrKind(modelFieldTypes, fields);
    expect(hasCorrectFieldTypes).toBeTruthy();
  });
});
