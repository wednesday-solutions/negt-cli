package modelUtils

import (
	"fmt"

	"github.com/wednesday-solutions/negt/gqlgenUtils/fileUtils"
)

// CreateGqlModelFiles will create GraphQl model files.
func CreateGqlModelFiles(modelName, dirName string, files, testFiles []string) error {

	path := fileUtils.FindDirectory(dirName)
	path = fmt.Sprintf("%s/%s", path, dirName)

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

// CreateCustomResolverFiles will create GraphQl model custom resolver files.
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
