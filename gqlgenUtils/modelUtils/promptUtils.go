package modelUtils

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/wednesday-solutions/negt/gqlgenUtils/fileUtils"
)

func PromptValidate(input string) error {
	if len(input) <= 2 {
		return errors.New("Invalid input.")
	}
	return nil
}

func PromptGetInput(pc PromptContent) string {
	validate := PromptValidate
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }}",
		Valid:   "{{ . | green }}",
		Invalid: "{{ . | red }}",
		Success: "{{ . | bold }}",
	}
	prompt := promptui.Prompt{
		Label:     pc.Label,
		Templates: templates,
		Validate:  validate,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt running failed %v\n", err)
		os.Exit(1)
	}
	return result
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

func PromptGetSelect(pc PromptContent) string {
	items := []string{"ID", "Int", "Float", "String", "Boolean", "DateTime"}
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.Label,
			Items:    items,
			AddLabel: "Other",
		}
		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return result
}

func PromptGetSelectPath(pc PromptContent) string {
	items := []string{"gql/models", "server/gql/models"}
	index := -1
	var result string
	var err error
	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.Label,
			Items:    items,
			AddLabel: "Other",
		}
		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}
	if result == "gql/models" || result == "server/gql/models" {
		status := fileUtils.DirExists(result)
		if !status {
			fmt.Println("gql/models directory is not exists, do 'negt gqlgen init'")
			os.Exit(1)
		}
	} else {
		status := fileUtils.DirExists(result)
		if !status {
			directories := strings.Split(result, "/")
			path := fileUtils.CurrentDirectory()
			for _, dir := range directories {
				if !fileUtils.IsExists(path, dir) {
					fileUtils.MakeDirectory(path, dir) // nolint:errcheck
				}
				path = fmt.Sprintf("%s/%s", path, dir)
			}
		}
	}
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return result
}
