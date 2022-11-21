package modelUtils

import (
	"fmt"
	"path/filepath"

	"github.com/ijasMohamad/cliApp/hbs"
)

func WriteTestCustomResolvers(
	modelName, dirName string, 
	fields, fieldTypes, resolverTestFiles []string, 
	nullFields []bool, 
	customMutation bool) error {

	path, _ := filepath.Abs(".")
	path = fmt.Sprintf("%s/%s/%s/tests", path, dirName, modelName)

	ctx := FieldUtils(modelName, fields, fieldTypes, nullFields, customMutation)

	for _, file := range resolverTestFiles {

		if file == "customCreateMutation.test.js" {
			err := hbs.CustomCreateMutationTestSource(modelName, path, file, ctx)
			if err != nil {
				return err
			}

		} else if file == "customUpdateMutation.test.js" {
			err := hbs.CustomUpdateMutationTestSource(modelName, path, file, ctx)
			if err != nil {
				return err
			}

		} else if file == "customDeleteMutation.test.js" {
			err := hbs.CustomDeleteMutationTestSource(modelName, path, file, ctx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
