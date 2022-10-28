package gqlgenUtils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	pluralize "github.com/gertd/go-pluralize"
	"github.com/aymerick/raymond"
)

func WriteModelFiles(modelName string, fields, fieldTypes, files []string) error {

	path, _ := filepath.Abs(".")
	path = fmt.Sprintf("%s/gql/models/%s", path, modelName)

	pluralize := pluralize.NewClient()

	lowerCaseModel := strings.ToLower(modelName)
	singularModel := pluralize.Singular(lowerCaseModel)
	pluralModel := pluralize.Plural(lowerCaseModel)
	titleSingularModel := strings.Title(singularModel)
	titlePluralModel := strings.Title(pluralModel)

	var intFlag, stringFlag bool
	var graphqlInt, graphqlString string

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
	if intFlag {
		graphqlInt = "GraphQLInt"
	}
	if stringFlag {
		graphqlString = "GraphQLString"
	}

	ctx := map[string]interface{}{
		"singularModel": singularModel,
		"pluralModel": pluralModel,
		"titleSingularModel": titleSingularModel,
		"titlePluralModel": titlePluralModel,
		"graphqlInt": graphqlInt,
		"graphqlString": graphqlString,
		"fields": fields,
		"fieldTypes": fieldTypes,
	}

	var fieldsWithType []string
	raymond.RegisterHelper("fieldsWithType", func(fields, fieldTypes []string) []string{
	
		if len(fieldsWithType) == 0 {
			for idx, field := range fields {
				fieldsWithType = append(fieldsWithType, fmt.Sprintf(`,
	%s: { type: %s }`, field, fieldTypes[idx]))
			}
		}
		return fieldsWithType
	})

	for _, file := range files {

	// Writing into index.js file.
		if file == "index.js" {

			source := `import { GraphQL{{titleSingularModel}} } from './model';
import { {{singularModel}}Connection } from './list';
import { {{titleSingularModel}}Queries } from './query';
import { {{singularModel}}Mutation } from './mutation';

// exporting graphql model of {{singularModel}}.
export const {{titleSingularModel}} = GraphQL{{titleSingularModel}};

// exporting {{singularModel}} connection
export const {{titleSingularModel}}Connection = {{singularModel}}Connection;

// exporting {{singularModel}} queries
export const {{singularModel}}Queries = {{titleSingularModel}}Queries;

// exporting mutations of {{singularModel}}
export const {{singularModel}}Mutations = {{singularModel}}Mutation;
`
			tpl, err := raymond.Parse(source)
			if err != nil {
				return err
			}
			result, err := tpl.Exec(ctx)
			if err != nil {
				return err
			}
			// Opening file by READ & WRITE permission.
			openFile, err := os.OpenFile(fmt.Sprintf("%s/%s", path, file), os.O_RDWR, 0644)
			if err != nil {
				return err
			}
			defer openFile.Close()

			// Write some text line-by-line to openFile.
			_, err = openFile.WriteString(result)
			if err != nil {
				return err
			}
			err = openFile.Sync()
			if err != nil {
				return err
			}
			fmt.Printf("%s file updated successfully. \n", file)
		}

		// Writing into model.js file.
		if file == "model.js" {

			source := `import { getNode }  from '@server/gql/node';
import { GraphQLID, GraphQLNonNull, GraphQLObjectType{{#if graphqlInt}}, {{graphqlInt}}{{/if}}{{#if graphqlString}}, {{graphqlString}}{{/if}} } from 'graphql';
import { getQueryFields, TYPE_ATTRIBUTES } from '@server/utils/gqlFieldUtils';

const { nodeInterface } = getNode();

export const {{singularModel}}Fields = {
	id: { type: new GraphQLNonNull(GraphQLID) }{{fieldsWithType fields fieldTypes}}
};

export const GraphQL{{titleSingularModel}} = new GraphQLObjectType({
	name: '{{titleSingularModel}}',
	interfaces: [nodeInterface],
	fields: () => ({
		...getQueryFields({{singularModel}}Fields, TYPE_ATTRIBUTES.isNonNull)
	})
});
`
			tpl, err := raymond.Parse(source)
			if err != nil {
				return err
			}
			result, err := tpl.Exec(ctx)
			if err != nil {
				return err
			}
			// Opening file by READ & WRITE permission.
			openFile, err := os.OpenFile(fmt.Sprintf("%s/%s", path, file), os.O_RDWR, 0644)
			if err != nil {
				return err
			}
			defer openFile.Close()

			_, err = openFile.WriteString(result)
			if err != nil {
				return err
			}
			err = openFile.Sync()
			if err != nil {
				return err
			}
			fmt.Printf("%s file updated successfully. \n", file)
		}

		// Writing into list.js file.
		if file == "list.js" {

			source := `import { createConnection } from 'graphql-sequelize';
import db from '@database/models';
import { GraphQL{{titleSingularModel}} } from './model';
import { sequelizedWhere } from '@server/database/dbUtils';
import { totalConnectionFields } from '@server/utils';

export const {{singularModel}}Connection = createConnection({
	name: '{{pluralModel}}',
	target: db.{{pluralModel}},
	nodeType: GraphQL{{titleSingularModel}},
	before: (findOptions, args, context) => {
		findOptions.include = findOptions.include || [];
		findOptions.where = sequelizedWhere(findOptions.where, args.where);
		return findOptions;
	},
	...totalConnectionFields
});
`
			tpl, err := raymond.Parse(source)
			if err != nil {
				return err
			}
			result, err := tpl.Exec(ctx)
			if err != nil {
				return err
			}
			openFile, err := os.OpenFile(fmt.Sprintf("%s/%s", path, file), os.O_RDWR, 0644)
			if err != nil {
				return err
			}
			defer openFile.Close()

			_, err = openFile.WriteString(result)
			if err != nil {
				return err
			}
			err = openFile.Sync()
			if err != nil {
				return err
			}
			fmt.Printf("%s file updated successfully.\n", file)
		}

		// writing into query.js file.
		if file == "query.js" {

			source := `import { GraphQLInt, GraphQLNonNull } from 'graphql';
import { GraphQL{{titleSingularModel}} } from './model';
import { {{singularModel}}Connection } from './list';
import db from '@database/models';

export const {{titleSingularModel}}Queries = {
	args: {
		id: {
			type: new GraphQLNonNull(GraphQLInt)
		}
	},
	query: {
		type: GraphQL{{titleSingularModel}}
	},
	list: {
		...{{singularModel}}Connection,
		resolve: {{singularModel}}Connection.resolve,
		type: {{singularModel}}Connection.connectionType,
		args: {{singularModel}}Connection.connectionArgs
	},
	model: db.{{pluralModel}}
};
`
			tpl, err := raymond.Parse(source)
			if err != nil {
				return err
			}
			result, err := tpl.Exec(ctx)
			if err != nil {
				return err
			}
			openFile, err := os.OpenFile(fmt.Sprintf("%s/%s", path, file), os.O_RDWR, 0644)
			if err != nil {
				return err
			}
			defer openFile.Close()

			_, err = openFile.WriteString(result)
			if err != nil {
				return err
			}
			err = openFile.Sync()
			if err != nil {
				return err
			}
			fmt.Printf("%s file updated successfully.\n", file)
		}

		// writing into mutation.js file.
		if file == "mutation.js" {

			source := `import { GraphQLID, GraphQLNonNull{{#if graphqlInt}}, {{graphqlInt}}{{/if}}{{#if graphqlString}}, {{graphqlString}}{{/if}} } from 'graphql';
import { GraphQL{{titleSingularModel}} } from './model';
import db from '@database/models';

export const {{singularModel}}MutationFields = {
	id: { type: new GraphQLNonNull(GraphQLID) }{{fieldsWithType fields fieldTypes}}
};

export const {{singularModel}}Mutation = {
	args: {{singularModel}}MutationFields,
	type: GraphQL{{titleSingularModel}},
	model: db.{{pluralModel}}
};
`
			tpl, err := raymond.Parse(source)
			if err != nil {
				return err
			}
			result, err := tpl.Exec(ctx)
			if err != nil {
				return err
			}
			openFile, err := os.OpenFile(fmt.Sprintf("%s/%s", path, file), os.O_RDWR, 0644)
			if err != nil {
				return err
			}
			defer openFile.Close()

			_, err = openFile.WriteString(result)
			if err != nil {
				return err
			}
			err = openFile.Sync()
			if err != nil {
				return err
			}
			fmt.Printf("%s file updated successfully.\n", file)
		}
	}
	return nil
}

func WriteModelTestFiles(modelName string, testFiles []string) error {

	return nil
}
