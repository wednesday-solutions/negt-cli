package modelUtils

import (
	"fmt"

	"github.com/ijasMohamad/negt/gqlgenUtils/fileUtils"
	"github.com/ijasMohamad/negt/hbs"
)

func WriteMockData(modelName, dirName string, fields, fieldTypes []string, nullFields []bool, customMutation bool) error {

	path := fileUtils.FindDirectory(dirName)

	dirName1 := "utils"
	dirName2 := "testUtils"
	file := fmt.Sprintf("%sMockData.js", modelName)

	if dirName == "server/gql/models" {
		path = fmt.Sprintf("%s/%s", path, "server")
		if !fileUtils.IsExists(path, dirName1) {
			_ = fileUtils.MakeDirectory(path, dirName1)
		}	
		path = fmt.Sprintf("%s/%s", path, dirName1)
		if !fileUtils.IsExists(path, dirName2) {
			_ = fileUtils.MakeDirectory(path, dirName2)
		}
		path = fmt.Sprintf("%s/%s", path, dirName2)

	} else {
		if !fileUtils.DirExists(dirName1) {
			_ = fileUtils.MakeDirectory(path, dirName1)
		}
		path = fmt.Sprintf("%s/%s", path, dirName1)
	
		if !fileUtils.DirExists(fmt.Sprintf("%s/%s", dirName1, dirName2)) {
			_ = fileUtils.MakeDirectory(path, dirName2)
		}
		path = fmt.Sprintf("%s/%s", path, dirName2)	
	}
	
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
