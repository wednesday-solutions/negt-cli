package fileUtils

import (
	"fmt"
	"os"
	"path/filepath"
)

func MakeDirectory(path string, dirName string) error {
	err := os.Mkdir(path+"/"+dirName, 0755)
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

func DirExists(dirName string) bool {
	path, _ := filepath.Abs(".")
	_, err := os.Stat(fmt.Sprintf("%s/%s", path, dirName))
	return err == nil
}

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
