package modelUtils

import (
	"fmt"
	"path/filepath"

	"github.com/ijasMohamad/cliApp/hbs"
)

func WriteCustomResolvers(
	modelName, dirName string, 
	fields, fieldTypes, resolverFiles []string, 
	nullFields []bool, 
	customMutation bool) error {

	path, _ := filepath.Abs(".")
	path = fmt.Sprintf("%s/%s/%s", path, dirName, modelName)

	ctx := FieldUtils(modelName, fields, fieldTypes, nullFields, customMutation)

	for _, file := range resolverFiles {

		if file == "customCreateMutation.js" {
			err := hbs.CustomCreateMutationSource(modelName, path, file, ctx)
			if err != nil {
				return err
			}

		} else if file == "customUpdateMutation.js" {
			err := hbs.CustomUpdateMutationSource(modelName, path, file, ctx)
			if err != nil {
				return err
			}

		} else if file == "customDeleteMutation.js" {
			err := hbs.CustomDeleteMutationSource(modelName, path, file, ctx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
