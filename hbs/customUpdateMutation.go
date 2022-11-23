package hbs

import "github.com/ijasMohamad/negt/gqlgenUtils/fileUtils"

func CustomUpdateMutationSource(modelName, path, file string, ctx map[string]interface{}) error {

	source := `export const customUpdateMutation = (model, args, context) => {
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
