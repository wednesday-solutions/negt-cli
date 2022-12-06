package fileUtils

import (
	"fmt"
	"os"
	"path/filepath"
)

// FindDirectory will give the path of an input.
func FindDirectory(dirName string) string {
	path, _ := filepath.Abs(".") // nolint:errcheck
	return path
}

// MakeDirectory will make directory according to input.
func MakeDirectory(path string, dirName string) error {
	err := os.Mkdir(path+"/"+dirName, 0755)
	if err != nil {
		return err
	}
	return nil
}

// MakeFile will create new file according to input path and file name.
func MakeFile(path string, fileName string) error {
	_, err := os.Create(path + "/" + fileName)
	if err != nil {
		return err
	}
	return nil
}

// DirExists will check the input directory is exist or not.
func DirExists(dirName string) bool {
	path, _ := filepath.Abs(".") // nolint:errcheck
	_, err := os.Stat(fmt.Sprintf("%s/%s", path, dirName))

	return err == nil
}

// IsExists will check the input directory is exists in the input path.
func IsExists(path, dirName string) bool {
	_, err := os.Stat(fmt.Sprintf("%s/%s", path, dirName))
	return err == nil
}

// WriteToFile for write input data into file with respect to path.
func WriteToFile(path, file, data string) error {
	// Opens file with read and write permission.
	openFile, err := os.OpenFile(fmt.Sprintf("%s/%s", path, file), os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error in openfile: ", err)
		return err
	}
	defer openFile.Close()

	_, err = openFile.WriteString(data)
	if err != nil {
		fmt.Println("Error in writeString: ", err)
		return err
	}
	err = openFile.Sync()
	if err != nil {
		fmt.Println("Error in sync: ", err)
		return err
	}
	fmt.Printf("%s file updated successfully. \n", file)

	return nil
}

// CurrentDirectory will give the root directory.
func CurrentDirectory() string {
	path, err := filepath.Abs(".")
	if err != nil {
		return ""
	}
	return path
}
