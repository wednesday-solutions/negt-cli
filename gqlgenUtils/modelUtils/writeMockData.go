package modelUtils

import (
	"fmt"
	"path/filepath"

	"github.com/ijasMohamad/cliApp/gqlgenUtils/fileUtils"
	"github.com/ijasMohamad/cliApp/hbs"
)

func WriteMockData(modelName string, fields, fieldTypes []string, nullFields []bool, customMutation bool) error {

	path, _ := filepath.Abs(".")
	dirName1 := "utils"
	dirName2 := "testUtils"
	file := fmt.Sprintf("%sMockData.js", modelName)

	if !fileUtils.DirExists(dirName1) {
		_ = fileUtils.MakeDirectory(path, dirName1)
	}
	path = fmt.Sprintf("%s/%s", path, dirName1)

	if !fileUtils.DirExists(fmt.Sprintf("%s/%s", dirName1, dirName2)) {
		_ = fileUtils.MakeDirectory(path, dirName2)
	}
	path = fmt.Sprintf("%s/%s", path, dirName2)

	err := fileUtils.MakeFile(path, file)
	if err != nil {
		return err
	}

	ctx :=  FieldUtils(modelName, fields, fieldTypes, nullFields, customMutation)

	err = hbs.MockDataSource(modelName, path, file, ctx)
	if err != nil {
		return err
	}

	return nil
}