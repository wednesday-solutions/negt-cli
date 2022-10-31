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
			for idx, field := range fields  {
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
