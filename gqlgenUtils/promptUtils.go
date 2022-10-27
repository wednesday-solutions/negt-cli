package gqlgenUtils

import (
	"errors"
	"fmt"
	"os"

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

func promptGetYesOrNoInput(pc promptContent) string {

	items := []string{ "Yes", "No"}
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
	return result
}

func promptGetSelect(pc promptContent) string {
	items := []string{"ID", "int", "string"}
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