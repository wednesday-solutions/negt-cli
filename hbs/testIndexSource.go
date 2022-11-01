package hbs

import (
	"github.com/ijasMohamad/cobra-cli/gqlgenUtils/fileUtils"
)

func IndexTestSource(modelName, path, file string, ctx map[string]interface{}) error {

	source := ``

	tpl, err := GenerateTemplate(source, ctx)
	if err != nil {
		return err
	}
	fileUtils.WriteToFile(path, file, tpl)

	return nil
}
