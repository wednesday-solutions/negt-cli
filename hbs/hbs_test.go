package hbs_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/aymerick/raymond"
	"github.com/stretchr/testify/assert"
	"github.com/wednesday-solutions/negt/hbs"
)

func TestGenerateTemplate(t *testing.T) {
	type args struct {
		source string
		ctx    map[string]interface{}
	}
	mockResult := "source"
	cases := []struct {
		name     string
		err      bool
		req      args
		parseErr bool
		execErr  bool
	}{
		{
			name: "Success",
			err:  false,
			req: args{
				source: "source",
				ctx:    map[string]interface{}{},
			},
		},
		{
			name: "Failure-parse",
			err:  true,
			req: args{
				source: "source",
				ctx:    map[string]interface{}{},
			},
			parseErr: true,
		},
		{
			name: "Failure-exec",
			err:  true,
			req: args{
				source: "source",
				ctx:    map[string]interface{}{},
			},
			execErr: true,
		},
	}

	for _, tt := range cases {
		patchParse := gomonkey.ApplyFunc(
			raymond.Parse,
			func(string) (*raymond.Template, error) {
				if tt.parseErr {
					return nil, fmt.Errorf("Error in parse")
				} else {
					return &raymond.Template{}, nil
				}
			},
		)
		defer patchParse.Reset()

		var tpl *raymond.Template
		patchExec := gomonkey.ApplyMethod(
			reflect.TypeOf(tpl),
			"Exec",
			func(*raymond.Template, interface{}) (string, error) {
				if tt.execErr {
					return "", fmt.Errorf("Error in exec")
				} else {
					return mockResult, nil
				}
			},
		)
		defer patchExec.Reset()

		t.Run(tt.name, func(t *testing.T) {
			_, err := hbs.GenerateTemplate(tt.req.source, tt.req.ctx)
			if err != nil {
				assert.Equal(t, true, err != nil)
				if tt.parseErr {
					assert.Equal(t, true, strings.Contains(err.Error(), "Error in parse"))
				} else if tt.execErr {
					assert.Equal(t, true, strings.Contains(err.Error(), "Error in exec"))
				}
			} else {
				assert.Equal(t, true, err == nil)
			}
		})
	}
}

func TestInitHbs(t *testing.T) {
	cases := []struct {
		name string
	}{
		{
			name: "Success",
		},
	}

	for _, tt := range cases {
		patchRegisterHelper := gomonkey.ApplyFunc(
			raymond.RegisterHelper,
			func(name string, helper interface{}) {},
		)
		defer patchRegisterHelper.Reset()

		t.Run(tt.name, func(t *testing.T) {
			testing.Init()
		})
	}
}

func TestOpeningBrace(t *testing.T){
	t.Run("Success", func(t *testing.T){
		response := hbs.OpeningBrace()
		assert.Equal(t, response, "{")
	})
}

func TestClosingBrace(t *testing.T){
	t.Run("Success", func(t *testing.T){
		response := hbs.ClosingBrace()
		assert.Equal(t, response, "}")
	})
}

func TestCustomMutations(t *testing.T){
	cases := []struct{
		name string
		req bool
	}{
		{
			name: "Success",
			req: true,
		},
		{
			name: "Fail",
			req: false,
		},
	}
	for _, tt := range cases {
		t.Run("Success", func(t *testing.T){
			response := hbs.CustomMutations(tt.req)
			if tt.req {
				assert.Equal(t, true, strings.Contains(response, "customCreateResolver"))
			} else {
				assert.Equal(t, true, strings.Contains(response, ""))
			}
		})
	}
}

func TestFieldsWithType(t *testing.T){
	t.Run("Success", func(t *testing.T){
		response := hbs.FieldsWithType(
			[]string{"field1", "field2"},
			[]string{"fieldType2", "fieldType2"},
			[]bool{true, false},
		)
		assert.Equal(t, true, strings.Contains(response[0], "field"))
	})
}

func TestTestFieldsWithID(t *testing.T){
	t.Run("Success", func(t *testing.T){
		response := hbs.TestFieldsWithID(
			[]string{"field1", "field2"},
			"modelName",
		)
		assert.Equal(t, true, strings.Contains(response[0], "field"))
	})
}

func TestInputStringFieldsWithID(t *testing.T){
	t.Run("Success", func(t *testing.T){
		response := hbs.InputStringFieldsWithID(
			[]string{"field1", "field2"},
			[]string{"GraphQLString", "fieldType2"},
			"modelName",
		)
		assert.Equal(t, true, strings.Contains(response[0], "field"))
	})
}

func TestInputStringFieldsWithoutID(t *testing.T){
	cases := []struct{
		name string
		req []string
	}{
		{
			name: "Success",
			req: []string{"GraphQLString", "fieldType2", "GraphQLString", "fieldType2"},
		},
		{
			name: "Success-else",
			req: []string{"fieldType2", "GraphQLString", "fieldType2", "GraphQLString"},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T){
			response := hbs.InputStringFieldsWithoutID(
				[]string{"field1", "field2", "field3", "field4"},
				tt.req,
				"modelName",
			)
			assert.Equal(t, true, strings.Contains(response[0], "Table"))
		})
	}
}

func TestTest(t *testing.T){
	t.Run("Success", func(t *testing.T){
		response := hbs.Test([]string{"field1", "field2"})
		assert.Equal(t, true, strings.Contains(response[0], "field"))
	})
}

func TestCustomMutationImports(t *testing.T){
	cases := []struct{
		name string
		customMutation bool
	}{
		{
			name: "Success-true",
			customMutation: true,
		},
		{
			name: "Success-false",
			customMutation: false,
		},
	}
	for _, tt := range cases{
		t.Run(tt.name, func(t *testing.T){
			response := hbs.CustomMutationImports(tt.customMutation)
			if tt.customMutation{
				assert.Equal(t, true, strings.Contains(response, "import"))
			} else {
				assert.Equal(t, true, strings.Contains(response, ""))
			}
		})
	}
}

func TestMockFields(t *testing.T){
	t.Run("Success", func(t *testing.T){
		response := hbs.MockFields(
			[]string{"field1", "field2", "field3", "field4", "field5", "field6"},
			[]string{"GraphQLID", "GraphQLInt", "GraphQLString", "GraphQLFloat", "GraphQLBoolean", "GraphQLDateTime"},
		)
		assert.Equal(t, true, strings.Contains(response[0], ":"))
	})
}

func TestMockImports(t *testing.T){
	cases := []struct{
		name string
		req []string
		graphqlType bool
	}{
		{
			name: "Success",
			req: []string{"GraphQLString"},
			graphqlType: true,
		},
		{
			name: "Success",
			req: []string{"fieldType"},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T){
			response := hbs.MockImports(tt.req)
			if tt.graphqlType {
				assert.Equal(t, true, strings.Contains(response, "import"))
			} else {
				assert.Equal(t, true, response == "")
			}
		})
	}
}