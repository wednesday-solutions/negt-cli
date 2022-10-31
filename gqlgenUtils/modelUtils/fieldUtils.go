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

	var graphqlID, graphqlInt, graphqlFloat, graphqlString, graphqlBoolean string
	var graphqlCustom string

	for idx, fieldType := range fieldTypes {
		if fieldType == "ID" {
			graphqlID = "GraphQLID"
			fieldTypes[idx] = graphqlID

		} else if fieldType == "Int" {
			graphqlInt = "GraphQLInt"
			fieldTypes[idx] = graphqlInt

		} else if fieldType == "Float" {
			graphqlFloat = "GraphQLFloat"
			fieldTypes[idx] = graphqlFloat

		} else if fieldType == "String" {
			graphqlString = "GraphQLString"
			fieldTypes[idx] = graphqlString

		} else if fieldType == "Boolean" {
			graphqlBoolean = "GraphQLBoolean"
			fieldTypes[idx] = graphqlBoolean

		} else {
			// From here we wan't to make relation of two graphql models
			graphqlCustom = string(fieldType[idx])
		}
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
		"graphqlCustom":      graphqlCustom,
		"fields":             fields,
		"fieldTypes":         fieldTypes,
		"nullFields":         nullFields,
	}
	return ctx
}
