package modelUtils_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/wednesday-solutions/negt/gqlgenUtils/modelUtils"
)

func TestCreateNewModel(t *testing.T) {

	cases := []struct {
		name           string
		err            bool
		YesOrNoInput   bool
		createGqlModelFilesErr bool
		writeModelFilesErr bool
		writeModelTestFilesErr bool
		writeMockDataErr bool
	}{
		{
			name:         "Success-No",
			err:          false,
			YesOrNoInput: false,
		},
		{
			name:         "Success-Yes",
			err:          false,
			YesOrNoInput: true,
		},
		{
			name:         "Fail-CreateGqlModelFiles",
			err:          true,
			createGqlModelFilesErr: true,
		},
		{
			name:         "Fail-WriteModelFiles",
			err:          true,
			writeModelFilesErr: true,
		},
		{
			name:         "Fail-WriteModelTestFiles",
			err:          true,
			writeModelTestFilesErr: true,
		},
		{
			name:         "Fail-WriteMockData",
			err:          true,
			writeMockDataErr: true,
		},
	}

	for _, tt := range cases {

		patchPromptGetSelectPath := gomonkey.ApplyFunc(
			modelUtils.PromptGetSelectPath,
			func(modelUtils.PromptContent) string {
				return "dirName"
			},
		)
		defer patchPromptGetSelectPath.Reset()

		patchPromptGetInput := gomonkey.ApplyFunc(
			modelUtils.PromptGetInput,
			func(modelUtils.PromptContent) string {
				return "modelName"
			},
		)
		defer patchPromptGetInput.Reset()

		patchPromptGetSelect := gomonkey.ApplyFunc(
			modelUtils.PromptGetSelect,
			func(modelUtils.PromptContent) string {
				return "field"
			},
		)
		defer patchPromptGetSelect.Reset()

		flag := true
		patchPromptGetYesOrNo := gomonkey.ApplyFunc(
			modelUtils.PromptGetYesOrNoInput,
			func(modelUtils.PromptContent) bool {
				if tt.YesOrNoInput {
					if flag {
						flag = false
						return true
					} else {
						return false
					}
				} else {
					return false
				}
			},
		)
		defer patchPromptGetYesOrNo.Reset()

		patchCreateGqlModelFiles := gomonkey.ApplyFunc(
			modelUtils.CreateGqlModelFiles,
			func(string, string, []string, []string) error {
				if tt.createGqlModelFilesErr {
					return fmt.Errorf("Error in CreateGqlModelFiles")
				} else {
					return nil
				}
			},
		)
		defer patchCreateGqlModelFiles.Reset()

		patchWriteModelFiles := gomonkey.ApplyFunc(
			modelUtils.WriteModelFiles,
			func(string, string, []string, []string, []string, []bool, bool) error {
				if tt.writeModelFilesErr {
					return fmt.Errorf("Error in WriteModelFiles")
				} else {
					return nil
				}
			},
		)
		defer patchWriteModelFiles.Reset()

		patchWriteModelTestFiles := gomonkey.ApplyFunc(
			modelUtils.WriteModelTestFiles,
			func(string, string, []string, []string, []string, []bool, bool) error {
				if tt.writeModelTestFilesErr {
					return fmt.Errorf("Error in WriteModelTestFiles")
				} else {
					return nil
				}
			},
		)
		defer patchWriteModelTestFiles.Reset()

		patchWriteMockData := gomonkey.ApplyFunc(
			modelUtils.WriteMockData,
			func(string, string, []string, []string, []bool, bool) error {
				if tt.writeMockDataErr {
					return fmt.Errorf("Error in WriteMockData")
				} else {
					return nil
				}
			},
		)
		defer patchWriteMockData.Reset()

		t.Run(tt.name, func(t *testing.T) {
			err := modelUtils.CreateNewModel()
			if err != nil {
				assert.Equal(t, true, err != nil)
				if tt.createGqlModelFilesErr{
					assert.Equal(t, true, strings.Contains(err.Error(), "Error in CreateGqlModelFiles"))
				} else if tt.writeModelFilesErr{
					assert.Equal(t, true, strings.Contains(err.Error(), "Error in WriteModelFiles"))
				} else if tt.writeModelTestFilesErr{
					assert.Equal(t, true, strings.Contains(err.Error(), "Error in WriteModelTestFiles"))
				} else if tt.writeMockDataErr{
					assert.Equal(t, true, strings.Contains(err.Error(), "Error in WriteMockData"))
				}
			} else {
				assert.Equal(t, true, err == nil)
			}
		})
	}
}

func TestAddField(t *testing.T){

	patchPromptGetInput := gomonkey.ApplyFunc(
		modelUtils.PromptGetInput,
		func(modelUtils.PromptContent) string {
			return "modelName"
		},
	)
	defer patchPromptGetInput.Reset()

	patchPromptGetSelect := gomonkey.ApplyFunc(
		modelUtils.PromptGetSelect,
		func(modelUtils.PromptContent) string {
			return "field"
		},
	)
	defer patchPromptGetSelect.Reset()

	patchPromptGetYesOrNo := gomonkey.ApplyFunc(
		modelUtils.PromptGetYesOrNoInput,
		func(modelUtils.PromptContent) bool {
			return false
		},
	)
	defer patchPromptGetYesOrNo.Reset()

	t.Run("Success", func(t *testing.T){
		fields, fieldTypes, nullFields :=  modelUtils.AddField(
			"modelName",
			[]string{"field"}, 
			[]string{"fieldType"},
			[]bool{true},
		)
		assert.Equal(t, fields[0], "field")
		assert.Equal(t, fieldTypes[0], "fieldType")
		assert.Equal(t, nullFields[0], true)
	})
}

func TestAddCustomMutations(t *testing.T){
	cases := []struct{
		name string
		createCustomResolverErr bool
		writeCustomResolverErr bool
		writeTestCustomResolverErr bool
	}{
		{
			name: "Success",
		},
		{
			name: "Fail-CreateCustomResolver",
			createCustomResolverErr: true,
		},
		{
			name: "Fail-WriteCustomResolver",
			writeCustomResolverErr: true,
		},
		{
			name: "Fail-CreateCustomResolver",
			writeTestCustomResolverErr: true,
		},
	}
	for _, tt := range cases {

		patchCreateCustomResolverFiles := gomonkey.ApplyFunc(
			modelUtils.CreateCustomResolverFiles,
			func(string, string, []string, []string) error {
				if tt.createCustomResolverErr{
					return fmt.Errorf("Error in CreateCustomResolverFiles")
				} else {
					return nil
				}
			},
		)
		defer patchCreateCustomResolverFiles.Reset()

		patchWriteCustomResolvers := gomonkey.ApplyFunc(
			modelUtils.WriteCustomResolvers,
			func(string, string, []string, []string, []string, []bool, bool) error {
				if tt.writeCustomResolverErr{
					return fmt.Errorf("Error in WriteCustomResolverFiles")
				} else {
					return nil
				}
			},
		)
		defer patchWriteCustomResolvers.Reset()

		patchWriteTestCustomResolvers := gomonkey.ApplyFunc(
			modelUtils.WriteTestCustomResolvers,
			func(string, string, []string, []string, []string, []bool, bool) error {
				if tt.writeTestCustomResolverErr{
					return fmt.Errorf("Error in WriteTestCustomResolverFiles")
				} else {
					return nil
				}
			},
		)
		defer patchWriteTestCustomResolvers.Reset()

		t.Run(tt.name, func(t *testing.T){
			err := modelUtils.AddCustomMutations(
				"modelName", 
				"dirName",
				[]string{"field"},
				[]string{"fieldType"},
				[]bool{true},
				true,
			)
			if err != nil {
				assert.Equal(t, true, err != nil)
				if tt.createCustomResolverErr{
					assert.Equal(t, true, strings.Contains(err.Error(), "Error in CreateCustomResolverFiles"))
				} else if tt.writeCustomResolverErr{
					assert.Equal(t, true, strings.Contains(err.Error(), "Error in WriteCustomResolverFiles"))
				} else if tt.writeTestCustomResolverErr{
					assert.Equal(t, true, strings.Contains(err.Error(), "Error in WriteTestCustomResolverFiles"))
				}
			} else {
				assert.Equal(t, true, err == nil)
			}
		})
	}
}