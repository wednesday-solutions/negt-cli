package gqlgenUtils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	pluralize "github.com/gertd/go-pluralize"
)

func WriteModelFiles(modelName string, fields, fieldTypes, files []string) error {

	path, _ := filepath.Abs(".")
	path = fmt.Sprintf("%s/gql/models/%s", path, modelName)

	pluralize := pluralize.NewClient()

	lowerCaseModel := strings.ToLower(modelName)
	singularModel := pluralize.Singular(lowerCaseModel)
	pluralModel := pluralize.Plural(lowerCaseModel)
	titleSingularModel := strings.Title(singularModel)
	// titlePluralModel := strings.Title(pluralModel)

	for _, file := range files {
		if file == "index.js" {
			// Opening file using READ & WRITE permission.
			openFile, err := os.OpenFile(fmt.Sprintf("%s/%s", path, file), os.O_RDWR, 0644)
			if err != nil {
				return err
			}
			defer openFile.Close()

			// Write some text line-by-line to openFile.
			_, err = openFile.WriteString(fmt.Sprintf(`import { GraphQL%s } from './model';
import { %sConnection } from './list';
import { %sQueries } from './query';
import { %sMutation } from './mutation';

// exporting graphql model of %s.
export const %s = GraphQL%s;

// exporting %s connection
export const %sConnection = %sConnection;

// exporting %s queries
export const %sQueries = %sQueries;

// exporting mutations of %s
export const %sMutations = %sMutation;

`, titleSingularModel,
singularModel,
titleSingularModel,
singularModel,
singularModel,
titleSingularModel,
titleSingularModel,
singularModel,
titleSingularModel,
singularModel,
singularModel,
singularModel,
titleSingularModel,
singularModel,
singularModel,
singularModel,
			))

			if err != nil {
				return err
			}
			err = openFile.Sync()
			if err != nil {
				return err
			}
			fmt.Printf("%s file updated successfully. \n", file)
		}

		if file == "model.js" {
			// Opening file using READ & WRITE permission.
			openFile, err := os.OpenFile(fmt.Sprintf("%s/%s", path, file), os.O_RDWR, 0644)
			if err != nil {
				return err
			}
			defer openFile.Close()

			var intFlag, stringFlag bool

			for idx, fieldType := range fieldTypes {
				if fieldType == "int" {
					fieldTypes[idx] = "GraphQLInt"
					intFlag = true
				}
				if fieldType == "string" {
					fieldTypes[idx] = "GraphQLString"
					stringFlag = true
				}
			}

			// Write some text line-by-line to openFile.
			_, err = openFile.WriteString(fmt.Sprintf(`import { getNode } from '@server/gql/node';
import { GraphQLID, GraphQLNonNull, GraphQLObjectType`))

			if intFlag {
				_, err = openFile.WriteString(`, GraphQLInt`)
			}
			if stringFlag {
				_, err = openFile.WriteString(`, GraphQLString`)
			}

			_, err = openFile.WriteString(fmt.Sprintf(` } from 'graphql';
import { getQueryFields, TYPE_ATTRIBUTES } from '@server/utils/gqlFieldUtils';

const { nodeInterface } = getNode();

export const %sFields = {
	id: { type: new GraphQLNonNull(GraphQLID) }`, singularModel))
			if err != nil {
				return err
			}

			for idx, field := range fields {

				_, err = openFile.WriteString(fmt.Sprintf(`,
	%s: { type: %s }`, field, fieldTypes[idx]))

				if err != nil {
					return err
				}
			}
			_, err = openFile.WriteString(fmt.Sprintf(`
};

export const GraphQL%s = new GraphQLObjectType({
	name: '%s',
	interfaces: [nodeInterface],
	fields: () => ({
		...getQueryFields(%sFields, TYPE_ATTRIBUTES.isNonNull)
	})
});
`,titleSingularModel, titleSingularModel, pluralModel))

			err = openFile.Sync()
			if err != nil {
				return err
			}
			fmt.Printf("%s file updated successfully. \n", file)
		}

		if file == "list.js" {

		}
		if file == "query.js" {

		}
		if file == "mutation.js" {

		}
	}

	return nil
}

func WriteModelTestFiles(modelName string, testFiles []string) error {

	return nil
}
