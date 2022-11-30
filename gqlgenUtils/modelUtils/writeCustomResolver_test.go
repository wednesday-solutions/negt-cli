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

func TestWriteCustomResolvers(t *testing.T) {

	type args struct {
		modelName      string
		dirName        string
		fields         []string
		fieldTypes     []string
		resolverFiles  []string
		nullFields     []bool
		customMutation bool
	}
	cases := []struct {
		name      string
		req       args
		createErr bool
		updateErr bool
		deleteErr bool
	}{
		{
			name: "Success",
			req: args{
				modelName:      "model",
				dirName:        "dir",
				fields:         []string{"field"},
				fieldTypes:     []string{"fieldType"},
				resolverFiles:  []string{"customCreateMutation.js", "customUpdateMutation.js", "customDeleteMutation.js"},
				nullFields:     []bool{true, false},
				customMutation: true,
			},
		},
		{
			name: "Fail-create",
			req: args{
				modelName:      "model",
				dirName:        "dir",
				fields:         []string{"field"},
				fieldTypes:     []string{"fieldType"},
				resolverFiles:  []string{"customCreateMutation.js"},
				nullFields:     []bool{true, false},
				customMutation: true,
			},
			createErr: true,
		},
		{
			name: "Fail-update",
			req: args{
				modelName:      "model",
				dirName:        "dir",
				fields:         []string{"field"},
				fieldTypes:     []string{"fieldType"},
				resolverFiles:  []string{"customUpdateMutation.js"},
				nullFields:     []bool{true, false},
				customMutation: true,
			},
			updateErr: true,
		},
		{
			name: "Fail-delete",
			req: args{
				modelName:      "model",
				dirName:        "dir",
				fields:         []string{"field"},
				fieldTypes:     []string{"fieldType"},
				resolverFiles:  []string{"customDeleteMutation.js"},
				nullFields:     []bool{true, false},
				customMutation: true,
			},
			deleteErr: true,
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

		patchCreateSource := gomonkey.ApplyFunc(
			hbs.CustomCreateMutationSource,
			func(string, string, string, map[string]interface{}) error {
				if tt.createErr {
					return fmt.Errorf("Error in createSource")
				} else {
					return nil
				}
			},
		)
		defer patchCreateSource.Reset()

		patchUpdateSource := gomonkey.ApplyFunc(
			hbs.CustomUpdateMutationSource,
			func(string, string, string, map[string]interface{}) error {
				if tt.updateErr {
					return fmt.Errorf("Error in updateSource")
				} else {
					return nil
				}
			},
		)
		defer patchUpdateSource.Reset()

		patchDeleteSource := gomonkey.ApplyFunc(
			hbs.CustomDeleteMutationSource,
			func(string, string, string, map[string]interface{}) error {
				if tt.deleteErr {
					return fmt.Errorf("Error in deleteSource")
				} else {
					return nil
				}
			},
		)
		defer patchDeleteSource.Reset()

		t.Run(tt.name, func(t *testing.T) {
			err := modelUtils.WriteCustomResolvers(
				tt.req.modelName,
				tt.req.dirName,
				tt.req.fields,
				tt.req.fieldTypes,
				tt.req.resolverFiles,
				tt.req.nullFields,
				tt.req.customMutation,
			)
			if err != nil {
				assert.Equal(t, true, err != nil)
				if tt.createErr {
					assert.Equal(t, true, strings.Contains(err.Error(), "Error in createSource"))
				} else if tt.updateErr {
					assert.Equal(t, true, strings.Contains(err.Error(), "Error in updateSource"))
				} else if tt.deleteErr {
					assert.Equal(t, true, strings.Contains(err.Error(), "Error in deleteSource"))
				}
			} else {
				assert.Equal(t, true, err == nil)
			}
		})
	}
}
