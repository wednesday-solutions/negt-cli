package gqlgenUtils

import (
	"os"
)

func MakeDirectory(path string, dirName string) error {
	err := os.Mkdir(path + "/" + dirName, 0755)
	if err != nil {
		return err
	}
	return nil
}

func MakeFile(path string, fileName string) error {
	_, err := os.Create(path + "/" + fileName)
	if err != nil {
		return err
	}
	return nil
}