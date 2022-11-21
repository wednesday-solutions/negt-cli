package modelUtils

import (
	pluralize "github.com/gertd/go-pluralize"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func FieldUtils(modelName string, fields, fieldTypes []string, nullFields []bool, customMutation bool) map[string]interface{} {

	pluralize := pluralize.NewClient()
	caser := cases.Title(language.English)

	singularModel := pluralize.Singular(modelName)
	pluralModel := pluralize.Plural(modelName)
	titleSingularModel := caser.String(singularModel)
	titlePluralModel := caser.String(pluralModel)

	var (
		graphqlID,
		graphqlInt,
		graphqlFloat,
		graphqlString,
		graphqlBoolean,
		graphqlDateTime,
		graphqlCustom string
	)

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

		} else if fieldType == "DateTime" {
			graphqlDateTime = "GraphQLDateTime"
			fieldTypes[idx] = graphqlDateTime

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
		"graphqlDateTime":    graphqlDateTime,
		"graphqlCustom":      graphqlCustom,
		"fields":             fields,
		"fieldTypes":         fieldTypes,
		"nullFields":         nullFields,
		"customMutation":     customMutation,
	}
	return ctx
}
