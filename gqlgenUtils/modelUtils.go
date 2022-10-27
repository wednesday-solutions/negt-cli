package gqlgenUtils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
)

type promptContent struct {
	errorMsg string
	label    string
}

func CreateNewModel() {
	modelNamePromptContent := promptContent{
		"What name would you like to make a model of? ",
		"Enter model name? ",
	}
	modelName := promptGetInput(modelNamePromptContent)

	filedPromptContent := promptContent{
		fmt.Sprintf("Which is the field would you like to add to %s model? ", modelName),
		fmt.Sprintf("Please provide a field for your model %s? ", modelName),
	}
	field := promptGetInput(filedPromptContent)

	typePromptContent := promptContent{
		fmt.Sprintf("Please provide the type for %s ", field),
		fmt.Sprintf("What is the type of the field %s? ", field),
	}
	FieldType := promptGetSelect(typePromptContent)

	fmt.Printf("Model Name: %s\n", modelName)
	fmt.Printf("Field: %s\n", field)
	fmt.Printf("Field Type: %s\n", FieldType)

	err := MakeGqlModel(modelName)
	if err != nil {
		fmt.Printf("%s already exists!", modelName)
		os.Exit(1)
	}

	err = WriteModelFiles()
	if err != nil {
		fmt.Printf("%s already exists!", modelName)
		os.Exit(1)
	}

	fmt.Printf("New GraphQL model '%s' created", modelName)
}

func promptGetInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
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

func MakeGqlModel(modelName string) error {

	path, _ := filepath.Abs(".")
	path = fmt.Sprintf("%s/gql/models", path)
	err := MakeDirectory(path, modelName)
	if err != nil {
		return err
	}
	Files := []string{
		"index.js",
		"model.js",
		"query.js",
		"list.js",
		"mutation.js",
	}
	path = fmt.Sprintf("%s/%s", path, modelName)
	for _, file := range Files {
		err := MakeFile(path, file)
		if err != nil {
			return err
		}
	}
	err = MakeDirectory(path, "tests")
	if err != nil {
		return err
	}
	TestFiles := []string{
		"index.test.js",
		"mutation.test.js",
		"query.test.js",
		"pagination.test.js",
	}
	path = fmt.Sprintf("%s/%s", path, "tests")
	for _, file := range TestFiles {
		err := MakeFile(path, file)
		if err != nil {
			return err
		}
	}
	return nil
}

func WriteModelFiles() error {

	return nil
}
