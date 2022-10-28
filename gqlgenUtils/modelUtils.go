package gqlgenUtils

import (
	"fmt"
	"os"
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
		fmt.Sprintf("Please provide field for your model %s? ", modelName),
	}
	field := promptGetInput(filedPromptContent)

	typePromptContent := promptContent{
		fmt.Sprintf("Please provide the type for %s ", field),
		fmt.Sprintf("What is the type of the field %s? ", field),
	}
	fieldType := promptGetSelect(typePromptContent)

	yesOrNoPromptContent := promptContent{
		fmt.Sprint("Do you want to add more fields to your model? "),
		fmt.Sprint("Do you want to add more fields? "),
	}
	var yesOrNo string
	var fields []string
	var fieldTypes []string
	fields = append(fields, field)
	fieldTypes = append(fieldTypes, fieldType)
	
	for yesOrNo != "No" {
		yesOrNo = promptGetYesOrNoInput(yesOrNoPromptContent)
		if yesOrNo == "Yes" {
			filedPromptContent := promptContent{
				fmt.Sprintf("Which is the another field would you like to add to %s model? ", modelName),
				fmt.Sprintf("Please provide another field for model %s? ", modelName),
			}
			field := promptGetInput(filedPromptContent)
			
			typePromptContent := promptContent{
				fmt.Sprintf("Please provide the type for %s ", field),
				fmt.Sprintf("What is the type of the field %s? ", field),
			}
			fieldType := promptGetSelect(typePromptContent)

			fields = append(fields, field)
			fieldTypes = append(fieldTypes, fieldType)
		}
	}

	fmt.Printf("Model Name: %s\n", modelName)
	fmt.Printf("Fields: %s\n", fields)
	fmt.Printf("Field Types: %s\n", fieldTypes)

	files := []string{
		"index.js",
		"model.js",
		"query.js",
		"list.js",
		"mutation.js",
	}
	testFiles := []string{
		"index.test.js",
		"mutation.test.js",
		"query.test.js",
		"pagination.test.js",
	}

	err := CreateGqlModel(modelName, files, testFiles)
	if err != nil {
		fmt.Printf("Error while creating files, %s", err)
		os.Exit(1)
	}

	err = WriteModelFiles(modelName, fields, fieldTypes, files)
	if err != nil {
		fmt.Printf("Error while writing into files, %s", err)
		os.Exit(1)
	}
	err = WriteModelTestFiles(modelName, testFiles)
	if err != nil {
		fmt.Printf("Error while writing into test files, %s", err)
		os.Exit(1)
	}

	fmt.Printf("New GraphQL model %s created!", modelName)
}
