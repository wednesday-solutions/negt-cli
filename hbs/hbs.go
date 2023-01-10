package hbs

import (
	"fmt"

	"github.com/aymerick/raymond"
)

func init() {
	raymond.RegisterHelper("fieldsWithType", FieldsWithType)
	raymond.RegisterHelper("testFieldsWithID", TestFieldsWithID)
	raymond.RegisterHelper("openingBrace", OpeningBrace)
	raymond.RegisterHelper("closingBrace", ClosingBrace)
	raymond.RegisterHelper("inputStringFieldsWithID", InputStringFieldsWithID)
	raymond.RegisterHelper("inputStringFieldsWithoutID", InputStringFieldsWithoutID)
	raymond.RegisterHelper("test", Test)
	raymond.RegisterHelper("customMutationImports", CustomMutationImports)
	raymond.RegisterHelper("customMutations", CustomMutations)
	raymond.RegisterHelper("mockFields", MockFields)
	raymond.RegisterHelper("mockImports", MockImports)
}

var fieldsWithType []string

// FieldsWithType is a helper function of Hbs.
func FieldsWithType(fields, fieldTypes []string, nullFields []bool) []string {

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
}

var testFields []string

// TestFieldsWithID is a helper function of Hbs.
func TestFieldsWithID(fields []string, modelName string) []string {
	if len(testFields) == 0 {
		for _, field := range fields {
			testFields = append(testFields, fmt.Sprintf(`,
				%s: %sTable[0].%s`, field, modelName, field))
		}
	}
	return testFields
}

var stringFields []string

// InputStringFieldsWithID is a helper function of Hbs.
func InputStringFieldsWithID(fields, fieldTypes []string, modelName string) []string {

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
}

var stringFieldsWithoutID []string

// InputStringFieldsWithoutID is a helper function of Hbs.
func InputStringFieldsWithoutID(fields, fieldTypes []string, modelName string) []string {

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
}

// Test is a helper function of Hbs.
func Test(fields []string) []string {
	return fields
}

var customMutationImports string

// CustomMutationImports is a helper function of Hbs.
func CustomMutationImports(customMutation bool) string {
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
}

var mockFields []string

// MockFields is a helper function of Hbs.
func MockFields(fields, fieldTypes []string) []string {
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
				fieldType = `'2022-02-02'`
			}

			mockFields = append(mockFields, fmt.Sprintf(`,
	%s: %s`, field, fieldType))
		}
	}
	return mockFields
}

// MockImports is a helper function of Hbs.
func MockImports(fieldTypes []string) string {
	for _, fieldType := range fieldTypes {
		if fieldType == "GraphQLString" {
			mockImports := `import faker from 'faker';`
			return mockImports
		}
	}
	return ""
}

// OpeningBrace is a helper function of Hbs.
func OpeningBrace() string {
	return "{"
}

// ClosingBrace is a helper function of Hbs.
func ClosingBrace() string {
	return "}"
}

var customMutations string

// CustomMutations is a helper function of Hbs.
func CustomMutations(customMutation bool) string {
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
}

// GenerateTemplate is a helper function of Hbs.
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
