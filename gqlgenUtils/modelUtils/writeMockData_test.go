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

func TestWriteMockData(t *testing.T) {

	type args struct {
		modelName      string
		dirName        string
		fields         []string
		fieldTypes     []string
		nullFields     []bool
		customMutation bool
	}
	cases := []struct {
		name        string
		req         args
		IsExists    bool
		makeFileErr bool
		hbsErr      bool
	}{
		{
			name: "Success",
			req: args{
				modelName:      "model",
				dirName:        "dir",
				fields:         []string{"field"},
				fieldTypes:     []string{"fieldType"},
				nullFields:     []bool{true, false},
				customMutation: true,
			},
			IsExists: true,
		},
		{
			name: "Success-negt",
			req: args{
				modelName:      "model",
				dirName:        "server/gql/models",
				fields:         []string{"field"},
				fieldTypes:     []string{"fieldType"},
				nullFields:     []bool{true, false},
				customMutation: true,
			},
			IsExists: false,
		},
		{
			name: "Fail-makeFile",
			req: args{
				modelName:      "model",
				dirName:        "dir",
				fields:         []string{"field"},
				fieldTypes:     []string{"fieldType"},
				nullFields:     []bool{true, false},
				customMutation: true,
			},
			IsExists:    true,
			makeFileErr: true,
		},
		{
			name: "Fail-mockData",
			req: args{
				modelName:      "model",
				dirName:        "dir",
				fields:         []string{"field"},
				fieldTypes:     []string{"fieldType"},
				nullFields:     []bool{true, false},
				customMutation: true,
			},
			IsExists: true,
			hbsErr:   true,
		},
	}

	for _, tt := range cases {
		patchFindDir := gomonkey.ApplyFunc(
			fileUtils.FindDirectory,
			func(string) string {
				return tt.req.dirName
			},
		)
		defer patchFindDir.Reset()

		patchIsExists := gomonkey.ApplyFunc(
			fileUtils.IsExists,
			func(string, string) bool {
				if tt.IsExists {
					return true
				} else {
					return false
				}
			},
		)
		defer patchIsExists.Reset()

		patchMakeDirectory := gomonkey.ApplyFunc(
			fileUtils.MakeDirectory,
			func(string, string) error {
				return nil
			},
		)
		defer patchMakeDirectory.Reset()

		patchMakeFile := gomonkey.ApplyFunc(
			fileUtils.MakeFile,
			func(string, string) error {
				if tt.makeFileErr {
					return fmt.Errorf("Error in MakeFile")
				} else {
					return nil
				}
			},
		)
		defer patchMakeFile.Reset()

		patchFieldUtils := gomonkey.ApplyFunc(
			modelUtils.FieldUtils,
			func(string, []string, []string, []bool, bool) map[string]interface{} {
				return map[string]interface{}{}
			},
		)
		defer patchFieldUtils.Reset()

		patchMockDataSource := gomonkey.ApplyFunc(
			hbs.MockDataSource,
			func(string, string, string, map[string]interface{}) error {
				if tt.hbsErr {
					return fmt.Errorf("Error in MockDataSource")
				} else {
					return nil
				}
			},
		)
		defer patchMockDataSource.Reset()

		t.Run(tt.name, func(t *testing.T) {
			err := modelUtils.WriteMockData(
				tt.req.modelName,
				tt.req.dirName,
				tt.req.fields,
				tt.req.fieldTypes,
				tt.req.nullFields,
				tt.req.customMutation,
			)
			if err != nil {
				assert.Equal(t, true, err != nil)
				if tt.makeFileErr {
					assert.Equal(t, true, strings.Contains(err.Error(), "Error in MakeFile"))
				} else if tt.hbsErr {
					assert.Equal(t, true, strings.Contains(err.Error(), "Error in MockDataSource"))
				}
			} else {
				assert.Equal(t, true, err == nil)
			}
		})
	}
}
