package modelUtils

import (
	"fmt"
	"os"
	"strings"

	"github.com/ijasMohamad/negt/gqlgenUtils/fileUtils"
)

func CreateGqlModelFiles(modelName, dirName string, files, testFiles []string) error {

	path := fileUtils.FindDirectory(dirName)

	if dirName != "gql/models" && dirName != "server/gql/models"{

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

	path := fileUtils.FindDirectory(dirName)

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
