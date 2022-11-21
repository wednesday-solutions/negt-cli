package hbs

import "github.com/ijasMohamad/cliApp/gqlgenUtils/fileUtils"

func MockDataSource(modelName, path, file string, ctx map[string]interface{}) error {

	source := `import range from 'lodash/range';
import faker from 'faker';

export const {{pluralModel}}Table = range(1, 10).map((_, index) => ({
	id: (index + 1).toString(){{mockFields fields fieldTypes}}
}));
`
	tpl, err := GenerateTemplate(source, ctx)
	if err != nil {
		return err
	}
	fileUtils.WriteToFile(path, file, tpl)

	return nil
}