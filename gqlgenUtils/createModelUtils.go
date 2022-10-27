package gqlgenUtils

import (
	"fmt"
	"path/filepath"
)

func CreateGqlModel(modelName string, files, testFiles []string) error {

	path, _ := filepath.Abs(".")
	path = fmt.Sprintf("%s/gql/models", path)
	err := MakeDirectory(path, modelName)
	if err != nil {
		return err
	}

	path = fmt.Sprintf("%s/%s", path, modelName)
	for _, file := range files {
		err := MakeFile(path, file)
		if err != nil {
			return err
		}
	}

	err = MakeDirectory(path, "tests")
	if err != nil {
		return err
	}

	path = fmt.Sprintf("%s/%s", path, "tests")
	for _, file := range testFiles {
		err := MakeFile(path, file)
		if err != nil {
			return err
		}
	}
	return nil
}
