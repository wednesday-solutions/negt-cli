package modelUtils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ijasMohamad/cliApp/gqlgenUtils/fileUtils"
)

func CreateGqlModelFiles(modelName, dirName string, files, testFiles []string) error {

	path, _ := filepath.Abs(".")
	if dirName != "gql/models" {

		directories := strings.Split(dirName, "/")
		if len(directories) > 0 {

			for _, dir := range directories {

				_, err := os.Stat(fmt.Sprintf("%s/%s", path, dir))
				if os.IsNotExist(err) {
					fileUtils.MakeDirectory(path, dir)
				}
				path = fmt.Sprintf("%s/%s", path, dir)
			}

		} else {
			_, err := os.Stat(fmt.Sprintf("%s/%s", path, directories[0]))
			if os.IsNotExist(err) {
				fileUtils.MakeDirectory(path, dirName)
			}
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

func CreateCustomResolverFiles(modelName, dirName string, resolverFiles, resolverTestFiles []string) error {

	path, _ := filepath.Abs(".")
	path = fmt.Sprintf("%s/%s/%s", path, dirName, modelName)

	for _, file := range resolverFiles {
		err := fileUtils.MakeFile(path, file)
		if err != nil {
			return err
		}
	}
	path = fmt.Sprintf("%s/%s", path, "tests")
	for _, file := range resolverTestFiles {
		err := fileUtils.MakeFile(path, file)
		if err != nil {
			return err
		}
	}
	return nil
}
