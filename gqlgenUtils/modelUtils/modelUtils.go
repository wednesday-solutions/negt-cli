package modelUtils

import (
	"fmt"
	"os/exec"
	"strings"

	pluralize "github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"github.com/wednesday-solutions/negt/gqlgenUtils/fileUtils"
)

type PromptContent struct {
	ErrorMsg string
	Label    string
}

func CreateNewModel() error {
	dirNamePromptContent := PromptContent{
		"In which directory you want to create GraphQL model? ",
		"Select directory? (If you are using Node-Express-GraphQL-Template, select 'server/gql/models') ",
	}
	dirName := PromptGetSelectPath(dirNamePromptContent)

	modelNamePromptContent := PromptContent{
		"What is the name of your GraphQL Model? ",
		"Enter GraphQL Model name? ",
	}
	modelName := PromptGetInput(modelNamePromptContent)

	// Making modelName into lowercase and plural.
	modelName = strings.ToLower(modelName)
	modelName = strcase.ToLowerCamel(modelName)
	pluralize := pluralize.NewClient()
	modelName = pluralize.Plural(modelName)

	filedPromptContent := PromptContent{
		fmt.Sprintf("Which is the field would you like to add to '%s' model? ", modelName),
		fmt.Sprintf("Please provide field for '%s'? ", modelName),
	}
	field := PromptGetInput(filedPromptContent)

	typePromptContent := PromptContent{
		fmt.Sprintf("Please provide the type for '%s' ", field),
		fmt.Sprintf("What is the type of the field '%s'? ", field),
	}
	fieldType := PromptGetSelect(typePromptContent)

	nullabilitylPromptContent := PromptContent{
		fmt.Sprintf("If the '%s' field is Non-null, then select 'Yes', else select 'No'? ", field),
		fmt.Sprintf("Do you want to make this '%s' field as Non null? ", field),
	}
	nullField := PromptGetYesOrNoInput(nullabilitylPromptContent)

	yesOrNoPromptContent := PromptContent{
		"Do you want to add more fields to your model? ",
		"Do you want to add more fields? ",
	}

	var (
		fields,
		fieldTypes []string
		yesOrNo    bool
		nullFields []bool
	)
	fields = append(fields, field)
	fieldTypes = append(fieldTypes, fieldType)
	nullFields = append(nullFields, nullField)

	yesOrNo = PromptGetYesOrNoInput(yesOrNoPromptContent)
	fmt.Println("YesOrNo: ", yesOrNo)
	for yesOrNo {
		fields, fieldTypes, nullFields = AddField(
			modelName, fields, fieldTypes, nullFields,
		)
		yesOrNo = PromptGetYesOrNoInput(yesOrNoPromptContent)
	}

	customMutationPromptContent := PromptContent{
		fmt.Sprintf("Do you want to make custom resolvers for your '%s' model? ", modelName),
		fmt.Sprintf("Do you need custom resolvers for your new '%s' model? ", modelName),
	}
	customMutation := PromptGetYesOrNoInput(customMutationPromptContent)

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
		return err
	}

	err = WriteModelFiles(modelName, dirName, fields, fieldTypes, files, nullFields, customMutation)
	if err != nil {
		fmt.Printf("Error while writing into files, %s", err)
		return err
	}

	err = WriteModelTestFiles(modelName, dirName, fields, fieldTypes, testFiles, nullFields, customMutation)
	if err != nil {
		fmt.Printf("Error while writing into test files, %s", err)
		return err
	}

	if customMutation {
		err := AddCustomMutations(modelName, dirName, fields, fieldTypes, nullFields, customMutation)
		if err != nil {
			return err
		}
	}

	err = WriteMockData(modelName, dirName, fields, fieldTypes, nullFields, customMutation)
	if err != nil {
		fmt.Printf("Error while writing into test custom resolvers, %s", err)
		return err
	}

	path := fileUtils.CurrentDirectory()
	
	cmd := exec.Command("npx", "eslint", path, "--fix")
	data, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error while executing lint command: ", string(data), err)
	}
	fmt.Printf("New GraphQL model %s created!", modelName)
	return nil
}

func AddField(
	modelName string,
	fields, fieldTypes []string,
	nullFields []bool,
) ([]string, []string, []bool) {

	filedPromptContent := PromptContent{
		fmt.Sprintf("Which is the another field would you like to add to '%s' model? ", modelName),
		fmt.Sprintf("Please provide another field for '%s'? ", modelName),
	}
	field := PromptGetInput(filedPromptContent)

	typePromptContent := PromptContent{
		fmt.Sprintf("Please provide the type for '%s' ", field),
		fmt.Sprintf("What is the type of the field '%s'? ", field),
	}
	fieldType := PromptGetSelect(typePromptContent)

	nullabilitylPromptContent := PromptContent{
		fmt.Sprintf("If the '%s' field is Non-null, then select 'Yes', else select 'No'? ", field),
		fmt.Sprintf("Do you want to make this '%s' field as Non null? ", field),
	}
	nullField := PromptGetYesOrNoInput(nullabilitylPromptContent)

	fields = append(fields, field)
	fieldTypes = append(fieldTypes, fieldType)
	nullFields = append(nullFields, nullField)

	return fields, fieldTypes, nullFields
}

func AddCustomMutations(
	modelName, dirName string,
	fields, fieldTypes []string,
	nullFields []bool,
	customMutation bool,
) error {
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
		return err
	}
	err = WriteCustomResolvers(modelName, dirName, fields, fieldTypes, resolverFiles, nullFields, customMutation)
	if err != nil {
		fmt.Printf("Error while writing into custom resolvers, %s", err)
		return err
	}
	err = WriteTestCustomResolvers(modelName, dirName, fields, fieldTypes, resolverTestFiles, nullFields, customMutation)
	if err != nil {
		fmt.Printf("Error while writing into test custom resolvers, %s", err)
		return err
	}
	return nil
}