package modelUtils

import (
	"fmt"
	"path/filepath"

	"github.com/ijasMohamad/cliApp/hbs"
)

func WriteModelTestFiles(
	modelName, dirName string, 
	fields, fieldTypes, testFiles []string, 
	nullFields []bool,
	customMutation bool) error {

	path, _ := filepath.Abs(".")
	path = fmt.Sprintf("%s/%s/%s/tests", path, dirName, modelName)

	ctx := FieldUtils(modelName, fields, fieldTypes, nullFields, customMutation)

	for _, file := range testFiles {

		if file == "index.test.js" {

			err := hbs.IndexTestSource(modelName, path, file, ctx)
			if err != nil {
				return err
			}

		} else if file == "model.test.js" {

			err := hbs.ModelTestSource(modelName, path, file, ctx)
			if err != nil {
				return err
			}

		} else if file == "query.test.js" {

			err := hbs.QueryTestSource(modelName, path, file, ctx)
			if err != nil {
				return err
			}

		} else if file == "list.test.js" {

			err := hbs.ListTestSource(modelName, path, file, ctx)
			if err != nil {
				return err
			}

		} else if file == "pagination.test.js" {

			err := hbs.PaginationTestSource(modelName, path, file, ctx)
			if err != nil {
				return err
			}

		} else if file == "mutation.test.js" {

			err := hbs.MutationTestSource(modelName, path, file, ctx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
