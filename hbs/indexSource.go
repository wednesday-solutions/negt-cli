package hbs

import "github.com/ijasMohamad/cliApp/gqlgenUtils/fileUtils"

func IndexSource(modelName, path, file string, ctx map[string]interface{}) error {

	source := `import { GraphQL{{titleSingularModel}} } from './model';
import { {{singularModel}}List } from './list';
import { {{singularModel}}Query } from './query';
import { {{singularModel}}Mutation } from './mutation';

const {{titleSingularModel}} = GraphQL{{titleSingularModel}};
const {{singularModel}}Lists = {{singularModel}}List;
const {{singularModel}}Queries = {{singularModel}}Query;
const {{singularModel}}Mutations = {{singularModel}}Mutation;

export { {{titleSingularModel}}, {{singularModel}}Lists, {{singularModel}}Queries, {{singularModel}}Mutations };
`
	tpl, err := GenerateTemplate(source, ctx)
	if err != nil {
		return err
	}
	fileUtils.WriteToFile(path, file, tpl)

	return nil
}
