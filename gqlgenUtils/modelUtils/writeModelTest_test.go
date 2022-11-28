package modelUtils_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/wednesday-solutions/negt/gqlgenUtils/fileUtils"
	"github.com/wednesday-solutions/negt/gqlgenUtils/modelUtils"
	"github.com/wednesday-solutions/negt/hbs"
)

func TestWriteModelTestFiles(t *testing.T) {

	type args struct {
		modelName      string
		dirName        string
		fields         []string
		fieldTypes     []string
		files          []string
		nullFields     []bool
		customMutation bool
	}
	cases := []struct {
		name          string
		req           args
		indexErr      bool
		modelErr      bool
		listErr       bool
		queryErr      bool
		mutationErr   bool
		paginationErr bool
	}{
		{
			name: "Success",
			req: args{
				modelName:      "model",
				dirName:        "dir",
				fields:         []string{"field"},
				fieldTypes:     []string{"fieldType"},
				files:          []string{"index.test.js", "model.test.js", "list.test.js", "query.test.js", "mutation.test.js", "pagination.test.js"},
				nullFields:     []bool{true, false},
				customMutation: true,
			},
		},
		{
			name: "Fail-index",
			req: args{
				modelName:      "model",
				dirName:        "dir",
				fields:         []string{"field"},
				fieldTypes:     []string{"fieldType"},
				files:          []string{"index.test.js"},
				nullFields:     []bool{true, false},
				customMutation: true,
			},
			indexErr: true,
		},
		{
			name: "Fail-model",
			req: args{
				modelName:      "model",
				dirName:        "dir",
				fields:         []string{"field"},
				fieldTypes:     []string{"fieldType"},
				files:          []string{"model.test.js"},
				nullFields:     []bool{true, false},
				customMutation: true,
			},
			modelErr: true,
		},
		{
			name: "Fail-list",
			req: args{
				modelName:      "model",
				dirName:        "dir",
				fields:         []string{"field"},
				fieldTypes:     []string{"fieldType"},
				files:          []string{"list.test.js"},
				nullFields:     []bool{true, false},
				customMutation: true,
			},
			listErr: true,
		},
		{
			name: "Fail-query",
			req: args{
				modelName:      "model",
				dirName:        "dir",
				fields:         []string{"field"},
				fieldTypes:     []string{"fieldType"},
				files:          []string{"query.test.js"},
				nullFields:     []bool{true, false},
				customMutation: true,
			},
			queryErr: true,
		},
		{
			name: "Fail-mutation",
			req: args{
				modelName:      "model",
				dirName:        "dir",
				fields:         []string{"field"},
				fieldTypes:     []string{"fieldType"},
				files:          []string{"mutation.test.js"},
				nullFields:     []bool{true, false},
				customMutation: true,
			},
			mutationErr: true,
		},
		{
			name: "Fail-pagination",
			req: args{
				modelName:      "model",
				dirName:        "dir",
				fields:         []string{"field"},
				fieldTypes:     []string{"fieldType"},
				files:          []string{"pagination.test.js"},
				nullFields:     []bool{true, false},
				customMutation: true,
			},
			paginationErr: true,
		},
	}

	for _, tt := range cases {
		patchFindDir := gomonkey.ApplyFunc(
			fileUtils.FindDirectory,
			func(string) string {
				return "path"
			},
		)
		defer patchFindDir.Reset()

		patchFieldUtils := gomonkey.ApplyFunc(
			modelUtils.FieldUtils,
			func(string, []string, []string, []bool, bool) map[string]interface{} {
				return map[string]interface{}{}
			},
		)
		defer patchFieldUtils.Reset()

		patchIndexSource := gomonkey.ApplyFunc(
			hbs.IndexTestSource,
			func(string, string, string, map[string]interface{}) error {
				if tt.indexErr {
					return fmt.Errorf("Error in indexSource")
				} else {
					return nil
				}
			},
		)
		defer patchIndexSource.Reset()

		patchModelSource := gomonkey.ApplyFunc(
			hbs.ModelTestSource,
			func(string, string, string, map[string]interface{}) error {
				if tt.modelErr {
					return fmt.Errorf("Error in modelSource")
				} else {
					return nil
				}
			},
		)
		defer patchModelSource.Reset()

		patchListSource := gomonkey.ApplyFunc(
			hbs.ListTestSource,
			func(string, string, string, map[string]interface{}) error {
				if tt.listErr {
					return fmt.Errorf("Error in listSource")
				} else {
					return nil
				}
			},
		)
		defer patchListSource.Reset()

		patchQuerySource := gomonkey.ApplyFunc(
			hbs.QueryTestSource,
			func(string, string, string, map[string]interface{}) error {
				if tt.queryErr {
					return fmt.Errorf("Error in querySource")
				} else {
					return nil
				}
			},
		)
		defer patchQuerySource.Reset()

		patchMutationSource := gomonkey.ApplyFunc(
			hbs.MutationTestSource,
			func(string, string, string, map[string]interface{}) error {
				if tt.mutationErr {
					return fmt.Errorf("Error in mutationSource")
				} else {
					return nil
				}
			},
		)
		defer patchMutationSource.Reset()

		patchPaginationSource := gomonkey.ApplyFunc(
			hbs.PaginationTestSource,
			func(string, string, string, map[string]interface{}) error {
				if tt.paginationErr {
					return fmt.Errorf("Error in paginationSource")
				} else {
					return nil
				}
			},
		)
		defer patchPaginationSource.Reset()

		t.Run(tt.name, func(t *testing.T) {
			err := modelUtils.WriteModelTestFiles(
				tt.req.modelName,
				tt.req.dirName,
				tt.req.fields,
				tt.req.fieldTypes,
				tt.req.files,
				tt.req.nullFields,
				tt.req.customMutation,
			)
			if err != nil {
				assert.Equal(t, true, err != nil)
				if tt.indexErr {
					assert.Equal(t, true, strings.Contains(err.Error(), "Error in indexSource"))
				} else if tt.modelErr {
					assert.Equal(t, true, strings.Contains(err.Error(), "Error in modelSource"))
				} else if tt.listErr {
					assert.Equal(t, true, strings.Contains(err.Error(), "Error in listSource"))
				} else if tt.queryErr {
					assert.Equal(t, true, strings.Contains(err.Error(), "Error in querySource"))
				} else if tt.mutationErr {
					assert.Equal(t, true, strings.Contains(err.Error(), "Error in mutationSource"))
				} else if tt.paginationErr {
					assert.Equal(t, true, strings.Contains(err.Error(), "Error in paginationSource"))
				}
			} else {
				assert.Equal(t, true, err == nil)
			}
		})
	}
}
