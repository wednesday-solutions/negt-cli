package modelUtils

import (
	"fmt"

	"github.com/wednesday-solutions/negt/gqlgenUtils/fileUtils"
	"github.com/wednesday-solutions/negt/hbs"
)

func WriteMockData(modelName, dirName string, fields, fieldTypes []string, nullFields []bool, customMutation bool) error {
	UTILS := "utils"
	TESTUTILS := "testUtils"
	file := fmt.Sprintf("%sMockData.js", modelName)

	path := fileUtils.FindDirectory(dirName)
	if dirName == "server/gql/models" {
		path = fmt.Sprintf("%s/%s", path, "server")
	}
	if !fileUtils.IsExists(path, UTILS) {
		_ = fileUtils.MakeDirectory(path, UTILS) // nolint:errcheck
	}
	path = fmt.Sprintf("%s/%s", path, UTILS)
	if !fileUtils.IsExists(path, TESTUTILS) {
		_ = fileUtils.MakeDirectory(path, TESTUTILS) // nolint:errcheck
	}
	path = fmt.Sprintf("%s/%s", path, TESTUTILS)

	err := fileUtils.MakeFile(path, file)
	if err != nil {
		return err
	}

	ctx := FieldUtils(modelName, fields, fieldTypes, nullFields, customMutation)

	err = hbs.MockDataSource(modelName, path, file, ctx)
	if err != nil {
		return err
	}
	return nil
}
