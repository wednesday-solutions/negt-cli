package modelUtils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	pluralize "github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
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
	modelName = strcase.ToLowerCamel(modelName)
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
		"Do you want to add more fields to your model? ",
		"Do you want to add more fields? ",
	}

	var (
		fields,
		fieldTypes []string
		yesOrNo    bool
		nullFields []bool
	)
	yesOrNo = true
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

	customMutationPromptContent := promptContent{
		fmt.Sprintf("Do you want to make custom resolvers for your %s model? ", modelName),
		fmt.Sprintf("Do you need custom resolvers for your new %s model? ", modelName),
	}
	customMutation := promptGetYesOrNoInput(customMutationPromptContent)

	files := []string{
		"index.js",
		"model.js",
		"query.js",
		"list.js",
		"mutation.js",
	}
	testFiles := []string{
		"index.test.js",
		"model.test.js",
		"query.test.js",
		"list.test.js",
		"pagination.test.js",
		"mutation.test.js",
	}

	err := CreateGqlModelFiles(modelName, dirName, files, testFiles)
	if err != nil {
		fmt.Printf("Error while creating files, %s", err)
		os.Exit(1)
	}

	err = WriteModelFiles(modelName, dirName, fields, fieldTypes, files, nullFields, customMutation)
	if err != nil {
		fmt.Printf("Error while writing into files, %s", err)
		os.Exit(1)
	}

	err = WriteModelTestFiles(modelName, dirName, fields, fieldTypes, testFiles, nullFields, customMutation)
	if err != nil {
		fmt.Printf("Error while writing into test files, %s", err)
		os.Exit(1)
	}

	if customMutation {
		resolverFiles := []string{
			"customCreateMutation.js",
			"customUpdateMutation.js",
			"customDeleteMutation.js",
		}
		resolverTestFiles := []string{
			"customCreateMutation.test.js",
			"customUpdateMutation.test.js",
			"customDeleteMutation.test.js",
		}
		err := CreateCustomResolverFiles(modelName, dirName, resolverFiles, resolverTestFiles)
		if err != nil {
			fmt.Printf("Error while creating files, %s", err)
			os.Exit(1)
		}
		err = WriteCustomResolvers(modelName, dirName, fields, fieldTypes, resolverFiles, nullFields, customMutation)
		if err != nil {
			fmt.Printf("Error while writing into custom resolvers, %s", err)
			os.Exit(1)
		}
		err = WriteTestCustomResolvers(modelName, dirName, fields, fieldTypes, resolverTestFiles, nullFields, customMutation)
		if err != nil {
			fmt.Printf("Error while writing into test custom resolvers, %s", err)
			os.Exit(1)
		}
	}

	err = WriteMockData(modelName, fields, fieldTypes, nullFields, customMutation)
	if err != nil {
		fmt.Printf("Error while writing into test custom resolvers, %s", err)
		os.Exit(1)
	}

	err = exec.Command("yarn", "lint").Run()
	if err != nil {
		fmt.Println("Error while executing script file", err)
	}

	fmt.Printf("New GraphQL model %s created!", modelName)
}
