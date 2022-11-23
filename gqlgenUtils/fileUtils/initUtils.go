package fileUtils

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type promptContent struct{
	errorMsg string
	label string
}

func InitUtils() bool {

	negtPromptContent := promptContent{
		"Are you using Node-Express-GraphQL-Template of Wednesday-solutions? ",
		"Are you using Node-Express-GraphQL-Template? ",
	}

	flag := promptGetYesOrNoInput(negtPromptContent)
	return flag
}


func promptGetYesOrNoInput(pc promptContent) bool {

	items := []string{"Yes", "No"}
	var index = -1
	var result string
	var err error
	prompt := promptui.Select{
		Label: pc.label,
		Items: items,
	}
	for index < 0 {
		index, result, err = prompt.Run()
		if err != nil {
			fmt.Println(pc.errorMsg)
		}
	}
	if result == "Yes" {
		return true
	} else {
		return false
	}
}