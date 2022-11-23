package hbs

import "github.com/wednesday-solutions/negt/gqlgenUtils/fileUtils"

func CustomCreateMutationSource(modelName, path, file string, ctx map[string]interface{}) error {

	source := `export const customCreateMutation = (model, args, context) => {
	return {}
}
`
	tpl, err := GenerateTemplate(source, ctx)
	if err != nil {
		return err
	}
	fileUtils.WriteToFile(path, file, tpl)

	return nil
}
