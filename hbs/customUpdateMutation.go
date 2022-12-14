package hbs

import "github.com/wednesday-solutions/negt/gqlgenUtils/fileUtils"

// CustomUpdateMutationSource is the source function for custom update mutation file.
func CustomUpdateMutationSource(modelName, path, file string, ctx map[string]interface{}) error {

	source := `export const customUpdateMutation = (model, args, context) => {
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
