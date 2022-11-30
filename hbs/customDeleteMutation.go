package hbs

import "github.com/wednesday-solutions/negt/gqlgenUtils/fileUtils"

func CustomDeleteMutationSource(modelName, path, file string, ctx map[string]interface{}) error {

	source := `export const customDeleteMutation = (model, args, context) => {
	return {}
}
`
	tpl, err := GenerateTemplate(source, ctx)
	if err != nil {
		return err
	}
	fileUtils.WriteToFile(path, file, tpl) // nolint:errcheck

	return nil
}
