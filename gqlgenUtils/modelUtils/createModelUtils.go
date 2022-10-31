package modelUtils

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/ijasMohamad/cobra-cli/gqlgenUtils/fileUtils"
)

func CreateGqlModelFiles(modelName, dirName string, files, testFiles []string) error {

	path, _ := filepath.Abs(".")
	if dirName != "gql/models" {
		directories := strings.Split(dirName, "/")
		if len(directories) > 0 {
			for _, dir := range directories {
				fileUtils.MakeDirectory(path, dir)
				path = fmt.Sprintf("%s/%s", path, dir)
			}
		} else {
			fileUtils.MakeDirectory(path, dirName)
			path = fmt.Sprintf("%s/%s", path, directories[0])
		}
	} else {
		path = fmt.Sprintf("%s/%s", path, dirName)
	}

	err := fileUtils.MakeDirectory(path, modelName)
	if err != nil {
		return err
	}

	path = fmt.Sprintf("%s/%s", path, modelName)
	for _, file := range files {
		err := fileUtils.MakeFile(path, file)
		if err != nil {
			return err
		}
	}

	err = fileUtils.MakeDirectory(path, "tests")
	if err != nil {
		return err
	}

	path = fmt.Sprintf("%s/%s", path, "tests")
	for _, file := range testFiles {
		err := fileUtils.MakeFile(path, file)
		if err != nil {
			return err
		}
	}
	return nil
}
