package modelUtils

import (
	"fmt"
	"path/filepath"

	"github.com/ijasMohamad/cliApp/hbs"
)

func WriteModelFiles(modelName, dirName string, fields, fieldTypes, files []string, nullFields []bool, customMutation bool) error {

	path, _ := filepath.Abs(".")
	path = fmt.Sprintf("%s/%s/%s", path, dirName, modelName)

	ctx := FieldUtils(modelName, fields, fieldTypes, nullFields, customMutation)

	for _, file := range files {

		// Writing into index.js file.
		if file == "index.js" {

			err := hbs.IndexSource(modelName, path, file, ctx)
			if err != nil {
				return err
			}

		// Writing into model.js file.
		} else if file == "model.js" {

			err := hbs.ModelSource(modelName, path, file, ctx)
			if err != nil {
				return err
			}

		// Writing into list.js file.
		} else if file == "list.js" {

			err := hbs.ListSource(modelName, path, file, ctx)
			if err != nil {
				return err
			}

		// writing into query.js file.
		} else if file == "query.js" {

			err := hbs.QuerySource(modelName, path, file, ctx)
			if err != nil {
				return err
			}

		// writing into mutation.js file.
		} else if file == "mutation.js" {

			err := hbs.MutationSource(modelName, path, file, ctx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
