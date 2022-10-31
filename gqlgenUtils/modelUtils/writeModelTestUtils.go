package modelUtils

import (
	"fmt"
	"path/filepath"

	"github.com/ijasMohamad/cobra-cli/hbs"
)

func WriteModelTestFiles(
	modelName, dirName string, 
	fields, fieldTypes, testFiles []string, 
	nullFields []bool) error {

	path, _ := filepath.Abs(".")
	path = fmt.Sprintf("%s/%s/%s/tests", path, dirName, modelName)

	ctx := FieldUtils(modelName, fields, fieldTypes, nullFields)

	for _, file := range testFiles {

		if file == fmt.Sprintf("%s.test.js", modelName) {

			err := hbs.ModelsTestSource(modelName, path, file, ctx)
			if err != nil {
				return err
			}
		}

		if file == "query.test.js" {

			err := hbs.QueryTestSource(modelName, path, file, ctx)
			if err != nil {
				return err
			}
		}

		if file == "pagination.test.js" {

			err := hbs.PaginationTestSource(modelName, path, file, ctx)
			if err != nil {
				return err
			}
		}

		if file == "mutation.test.js" {

			err := hbs.MutationTestSource(modelName, path, file, ctx)
			if err != nil {
				return err
			}
		}

	}

	return nil
}
