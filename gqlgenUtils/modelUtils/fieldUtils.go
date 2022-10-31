package modelUtils

import (
	"strings"

	pluralize "github.com/gertd/go-pluralize"
)

func FieldUtils(modelName string, fields, fieldTypes []string, nullFields []bool) map[string]interface{} {

	pluralize := pluralize.NewClient()

	singularModel := pluralize.Singular(modelName)
	pluralModel := pluralize.Plural(modelName)
	titleSingularModel := strings.Title(singularModel)
	titlePluralModel := strings.Title(pluralModel)

	var IDFlag, intFlag, floatFlag, stringFlag, booleanFlag bool
	var graphqlID, graphqlInt, graphqlFloat, graphqlString, graphqlBoolean string

	for idx, fieldType := range fieldTypes {
		if fieldType == "ID" {
			fieldTypes[idx] = "GraphQLID"
			IDFlag = true
		}
		if fieldType == "Int" {
			fieldTypes[idx] = "GraphQLInt"
			intFlag = true
		}
		if fieldType == "Float" {
			fieldTypes[idx] = "GraphQLFloat"
			floatFlag = true
		}
		if fieldType == "String" {
			fieldTypes[idx] = "GraphQLString"
			stringFlag = true
		}
		if fieldType == "Boolean" {
			fieldTypes[idx] = "GraphQLBoolean"
			booleanFlag = true
		}
	}
	if IDFlag {
		graphqlID = "GraphQLID"
	}
	if intFlag {
		graphqlInt = "GraphQLInt"
	}
	if floatFlag {
		graphqlFloat = "GraphQLFloat"
	}
	if stringFlag {
		graphqlString = "GraphQLString"
	}
	if booleanFlag {
		graphqlBoolean = "GraphQLBoolean"
	}

	ctx := map[string]interface{}{
		"singularModel":      singularModel,
		"pluralModel":        pluralModel,
		"titleSingularModel": titleSingularModel,
		"titlePluralModel":   titlePluralModel,
		"graphqlID":          graphqlID,
		"graphqlInt":         graphqlInt,
		"graphqlFloat":       graphqlFloat,
		"graphqlString":      graphqlString,
		"graphqlBoolean":     graphqlBoolean,
		"fields":             fields,
		"fieldTypes":         fieldTypes,
		"nullFields":         nullFields,
	}
	return ctx
}
