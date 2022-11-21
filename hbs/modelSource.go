package hbs

import "github.com/ijasMohamad/cliApp/gqlgenUtils/fileUtils"

func ModelSource(modelName, path, file string, ctx map[string]interface{})  error {

	source := `import { getNode }  from '@server/gql/node';
import { GraphQLID, GraphQLNonNull, GraphQLObjectType{{#if graphqlInt}}, {{graphqlInt}}{{/if}}{{#if graphqlString}}, {{graphqlString}}{{/if}}{{#if graphqlFloat}}, {{graphqlFloat}}{{/if}}{{#if graphqlBoolean}}, {{graphqlBoolean}}{{/if}}{{#if graphqlDateTime}}, {{graphqlDateTime}}{{/if}} } from 'graphql';
import { getQueryFields, TYPE_ATTRIBUTES } from '@server/utils/gqlFieldUtils';
import { timestamps } from '@gqlFields/timestamps';

const { nodeInterface } = getNode();

export const {{singularModel}}Fields = {
	id: { type: new GraphQLNonNull(GraphQLID) }{{fieldsWithType fields fieldTypes nullFields}}
};

export const GraphQL{{titleSingularModel}} = new GraphQLObjectType({
	name: '{{titleSingularModel}}',
	interfaces: [nodeInterface],
	fields: () => ({
		...getQueryFields({{singularModel}}Fields, TYPE_ATTRIBUTES.isNonNull),
		...timestamps
	})
});
`
	tpl, err := GenerateTemplate(source, ctx)
	if err != nil {
		return err
	}
	fileUtils.WriteToFile(path, file, tpl)

	return nil
}