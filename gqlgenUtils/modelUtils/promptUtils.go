package modelUtils

import (
	"errors"
	"fmt"
	"os"

	"github.com/ijasMohamad/cliApp/gqlgenUtils/fileUtils"
	"github.com/manifoldco/promptui"
)

func promptGetInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 1 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }}",
		Valid:   "{{ . | green }}",
		Invalid: "{{ . | red }}",
		Success: "{{ . | bold }}",
	}
	prompt := promptui.Prompt{
		Label:     pc.label,
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

func promptGetSelect(pc promptContent) string {
	items := []string{"ID", "Int", "Float", "String", "Boolean", "DateTime"}
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
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

func promptGetSelectPath(pc promptContent) string {
	items := []string{"gql/models"}
	index := -1
	var result string
	var err error
	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
			Items:    items,
			AddLabel: "Other",
		}
		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}
	if result == "gql/models" {
		status := fileUtils.DirExists(result)
		if !status {
			fmt.Println("gql/models directory is not exists, do 'negt gqlgen init'")
			os.Exit(1)
		}
	}
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return result
}
