package modelUtils

import (
	"fmt"

	"github.com/wednesday-solutions/negt/gqlgenUtils/fileUtils"
	"github.com/wednesday-solutions/negt/hbs"
)

// WriteCustomResolvers is the function for write data into custom resolver files.
func WriteCustomResolvers(
	modelName, dirName string,
	fields, fieldTypes, resolverFiles []string,
	nullFields []bool,
	customMutation bool) error {

	path := fileUtils.FindDirectory(dirName)

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
