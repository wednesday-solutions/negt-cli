package fileUtils

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type PromptContent struct {
	ErrorMsg string
	Label    string
}

func InitUtils() bool {

	negtPromptContent := PromptContent{
		"Are you using Node-Express-GraphQL-Template of Wednesday-solutions? ",
		"Are you using Node-Express-GraphQL-Template? ",
	}

	flag := PromptGetYesOrNoInput(negtPromptContent)
	return flag
}

func PromptGetYesOrNoInput(pc PromptContent) bool {

	items := []string{"Yes", "No"}
	var index = -1
	var result string
	var err error
	prompt := promptui.Select{
		Label: pc.Label,
		Items: items,
	}
	for index < 0 {
		index, result, err = prompt.Run()
		if err != nil {
			fmt.Println(pc.ErrorMsg)
		}
	}
	if result == "Yes" {
		return true
	} else {
		return false
	}
}
