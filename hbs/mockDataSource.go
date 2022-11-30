package hbs

import "github.com/wednesday-solutions/negt/gqlgenUtils/fileUtils"

func MockDataSource(modelName, path, file string, ctx map[string]interface{}) error {

	source := `import range from 'lodash/range';
{{{mockImports fieldTypes}}}

export const {{pluralModel}}Table = range(1, 10).map((_, index) => ({
	id: (index + 1).toString(){{mockFields fields fieldTypes}}
}));
`
	tpl, err := GenerateTemplate(source, ctx)
	if err != nil {
		return err
	}
	fileUtils.WriteToFile(path, file, tpl) // nolint:errcheck

	return nil
}
