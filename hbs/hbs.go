package hbs

import (
	"fmt"

	"github.com/aymerick/raymond"
)

func init() {

	var fieldsWithType []string
	raymond.RegisterHelper("fieldsWithType", func(fields, fieldTypes []string, nullFields []bool) []string {

		if len(fieldsWithType) == 0 {
			for idx, field := range fields {
				if nullFields[idx] {
					fieldsWithType = append(fieldsWithType, fmt.Sprintf(`,
	%s: { type: new GraphQLNonNull(%s) }`, field, fieldTypes[idx]))
				} else {
					fieldsWithType = append(fieldsWithType, fmt.Sprintf(`,
	%s: { type: %s }`, field, fieldTypes[idx]))
				}
			}
		}
		return fieldsWithType
	})

	var testFields []string
	raymond.RegisterHelper("testFieldsWithID", func(fields []string, modelName string) []string {
		if len(testFields) == 0 {
			for _, field := range fields {
				testFields = append(testFields, fmt.Sprintf(`,
				%s: %sTable[0].%s`, field, modelName, field))
			}
		}
		return testFields
	})

	raymond.RegisterHelper("openingBrace", func() string {
		return "{"
	})
	raymond.RegisterHelper("closingBrace", func() string {
		return "}"
	})

	var stringFields []string
	raymond.RegisterHelper("inputStringFieldsWithID", func(fields, fieldTypes []string, modelName string) []string {

		if len(stringFields) == 0 {
			for idx, field := range fields {
				if fieldTypes[idx] == "GraphQLString" {
					stringFields = append(stringFields, fmt.Sprintf(`,
			%s: "${%sTable[0].%s}"`, field, modelName, field))
				} else {
					stringFields = append(stringFields, fmt.Sprintf(`,
			%s: ${%sTable[0].%s}`, field, modelName, field))
				}
			}
		}
		return stringFields
	})

	var stringFieldsWithoutID []string
	raymond.RegisterHelper("inputStringFieldsWithoutID", func(fields, fieldTypes []string, modelName string) []string {

		if len(stringFieldsWithoutID) == 0 {
			for idx, field := range fields {
				if idx == 0 {
					if fieldTypes[idx] == "GraphQLString" {
						stringFieldsWithoutID = append(stringFieldsWithoutID, fmt.Sprintf(`
			%s: "${%sTable[0].%s}"`, field, modelName, field))
					} else {
						stringFieldsWithoutID = append(stringFieldsWithoutID, fmt.Sprintf(`
			%s: ${%sTable[0].%s}`, field, modelName, field))
					}
				} else {
					if fieldTypes[idx] == "GraphQLString" {
						stringFieldsWithoutID = append(stringFieldsWithoutID, fmt.Sprintf(`,
			%s: "${%sTable[0].%s}"`, field, modelName, field))
					} else {
						stringFieldsWithoutID = append(stringFieldsWithoutID, fmt.Sprintf(`,
			%s: ${%sTable[0].%s}`, field, modelName, field))
					}
				}
			}
		}
		return stringFieldsWithoutID
	})

	raymond.RegisterHelper("test", func(fields []string) []string {
		return fields
	})

	var customMutationImports string
	raymond.RegisterHelper("customMutationImports", func(customMutation bool) string {
		if customMutation {
			if customMutationImports == "" {
				customMutationImports = `
import { customCreateMutation } from './customCreateMutation';  
import { customUpdateMutation } from './customUpdateMutation';  
import { customDeleteMutation } from './customDeleteMutation';`
			}
			return customMutationImports
		}
		return ""
	})

	var customMutations string
	raymond.RegisterHelper("customMutations", func(customMutation bool) string {
		if customMutation {
			if customMutations == "" {
				customMutations = `,
	customCreateResolver: customCreateMutation,
	customUpdateResolver: customUpdateMutation,
	customDeleteResolver: customDeleteMutation`
			}
			return customMutations
		}
		return ""
	})

	var mockFields []string
	raymond.RegisterHelper("mockFields", func(fields, fieldTypes []string) []string {
		if len(mockFields) == 0 {
			for idx, field := range fields {
				var fieldType string
				if fieldTypes[idx] == "GraphQLID" {
					fieldType = "(index + 1)"
				} else if fieldTypes[idx] == "GraphQLInt" {
					fieldType = "10"
				} else if fieldTypes[idx] == "GraphQLString" {
					fieldType = "faker.name.firstName()"
				} else if fieldTypes[idx] == "GraphQLFloat" {
					fieldType = "15.54"
				} else if fieldTypes[idx] == "GraphQLBoolean" {
					fieldType = "true"
				} else if fieldTypes[idx] == "GraphQLDateTime" {
					fieldType = "2022-02-02"
				}

				mockFields = append(mockFields, fmt.Sprintf(`,
	%s: %s`, field, fieldType))
			}
		}
		return mockFields
	})
}

func GenerateTemplate(source string, ctx map[string]interface{}) (string, error) {

	tpl, err := raymond.Parse(source)
	if err != nil {
		fmt.Println("Error in gentpl: ", err)
		return "", err
	}
	result, err := tpl.Exec(ctx)
	if err != nil {
		fmt.Println("Error in gentpl Exec: ", err)
		return "", err
	}
	return result, nil
}
