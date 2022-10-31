package hbs

import "github.com/ijasMohamad/cobra-cli/gqlgenUtils/fileUtils"

func IndexSource (modelName, path, file string, ctx map[string]interface{}) error {
	
	source := `import { GraphQL{{titleSingularModel}} } from './model';
import { {{singularModel}}Connection } from './list';
import { {{titleSingularModel}}Queries } from './query';
import { {{singularModel}}Mutation } from './mutation';

const {{titleSingularModel}} = GraphQL{{titleSingularModel}};
const {{titleSingularModel}}Connection = {{singularModel}}Connection;
const {{singularModel}}Queries = {{titleSingularModel}}Queries;
const {{singularModel}}Mutations = {{singularModel}}Mutation;

export {
	{{titleSingularModel}},
	{{titleSingularModel}}Connection,
	{{singularModel}}Queries,
	{{singularModel}}Mutations
};
`
	tpl, err := GenerateTemplate(source, ctx)
	if err != nil {
		return err
	}
	fileUtils.WriteToFile(path, file, tpl)

	return nil
}
