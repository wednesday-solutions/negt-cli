package hbs

import "github.com/wednesday-solutions/negt/gqlgenUtils/fileUtils"

// CustomCreateMutationSource is the source function for custom create mutation file.
func CustomCreateMutationSource(modelName, path, file string, ctx map[string]interface{}) error {

	source := `export const customCreateMutation = (model, args, context) => {
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
