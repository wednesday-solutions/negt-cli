package modelUtils

import (
	"fmt"
	"os"
	"strings"

	pluralize "github.com/gertd/go-pluralize"
)

type promptContent struct {
	errorMsg string
	label    string
}

func CreateNewModel() {
	dirNamePromptContent := promptContent{
		"In which directory you want to create GraphQL model? ",
		"Select directory? ",
	}
	dirName := promptGetSelectPath(dirNamePromptContent)

	modelNamePromptContent := promptContent{
		"What is the name of your GraphQL Model? ",
		"Enter GraphQL Model name? ",
	}
	modelName := promptGetInput(modelNamePromptContent)

	// Making modelName into lowercase and plural.
	modelName = strings.ToLower(modelName)
	pluralize := pluralize.NewClient()
	modelName = pluralize.Plural(modelName)

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

	nullabilitylPromptContent := promptContent{
		fmt.Sprintf("If the %s field is Non-null, then select 'Yes', else select 'No'? ", field),
		fmt.Sprintf("Do you want to make this %s field as Non null? ", field),
	}
	nullField := promptGetYesOrNoInput(nullabilitylPromptContent)

	yesOrNoPromptContent := promptContent{
		fmt.Sprint("Do you want to add more fields to your model? "),
		fmt.Sprint("Do you want to add more fields? "),
	}

	yesOrNo := true
	var fields []string
	var fieldTypes []string
	var nullFields []bool
	fields = append(fields, field)
	fieldTypes = append(fieldTypes, fieldType)
	nullFields = append(nullFields, nullField)

	for yesOrNo {
		yesOrNo = promptGetYesOrNoInput(yesOrNoPromptContent)
		if yesOrNo {
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

			nullabilitylPromptContent := promptContent{
				fmt.Sprintf("If the %s field is Non-null, then select 'Yes', else select 'No'? ", field),
				fmt.Sprintf("Do you want to make this %s field as Non null? ", field),
			}
			nullField := promptGetYesOrNoInput(nullabilitylPromptContent)

			fields = append(fields, field)
			fieldTypes = append(fieldTypes, fieldType)
			nullFields = append(nullFields, nullField)
		}
	}

	files := []string{
		"index.js",
		"model.js",
		"query.js",
		"list.js",
		"mutation.js",
	}
	testFiles := []string{
		"query.test.js",
		"pagination.test.js",
		"mutation.test.js",
	}
	testFiles = append(testFiles, fmt.Sprintf("%s.test.js", modelName))

	err := CreateGqlModelFiles(modelName, dirName, files, testFiles)
	if err != nil {
		fmt.Printf("Error while creating files, %s", err)
		os.Exit(1)
	}

	err = WriteModelFiles(modelName, dirName, fields, fieldTypes, files, nullFields)
	if err != nil {
		fmt.Printf("Error while writing into files, %s", err)
		os.Exit(1)
	}
	
	err = WriteModelTestFiles(modelName, dirName, fields, fieldTypes, testFiles, nullFields)
	if err != nil {
		fmt.Printf("Error while writing into test files, %s", err)
		os.Exit(1)
	}

	fmt.Printf("New GraphQL model %s created!", modelName)
}
