package hbs

import (
	"github.com/ijasMohamad/cliApp/gqlgenUtils/fileUtils"
)

func IndexTestSource(modelName, path, file string, ctx map[string]interface{}) error {

	source := `import { GraphQL{{titleSingularModel}} } from './model';
import { {{singularModel}}List } from './list';
import { {{singularModel}}Query } from './query';
import { {{singularModel}}Mutation } from './mutation';
	
descibe('Checking importing is succesfull', () => {
	it('Should check imports', async () => {
		expect(GraphQL{{titleSingularModel}}).toBeTruthy();
		expect(Object.keys(GraphQL{{titleSingularModel}}).length > 0).toBe(true);
		expect({{singularModel}}List).toBeTruthy();
		expect(Object.keys({{singularModel}}List).length > 0).toBe(true);
		expect({{singularModel}}Query).toBeTruthy();
		expect(Object.keys({{singularModel}}Query).length > 0).toBe(true);
		expect({{singularModel}}Mutation).toBeTruthy();
		expect(Object.keys({{singularModel}}Mutation).length > 0).toBe(true);
	});
});
`

	tpl, err := GenerateTemplate(source, ctx)
	if err != nil {
		return err
	}
	fileUtils.WriteToFile(path, file, tpl)

	return nil
}
