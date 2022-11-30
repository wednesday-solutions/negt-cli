package fileUtils_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/manifoldco/promptui"
	"github.com/stretchr/testify/assert"
	"github.com/wednesday-solutions/negt/gqlgenUtils/fileUtils"
)

func TestInitUtils(t *testing.T) {

	cases := []struct {
		name string
		flag bool
	}{
		{
			name: "Success-true",
			flag: true,
		},
		{
			name: "Success-false",
			flag: false,
		},
	}

	for _, tt := range cases {

		patchPromptGetYesOrNo := gomonkey.ApplyFunc(
			fileUtils.PromptGetYesOrNoInput,
			func(fileUtils.PromptContent) bool {
				if tt.flag {
					return true
				} else {
					return false
				}
			},
		)
		defer patchPromptGetYesOrNo.Reset()

		t.Run(tt.name, func(t *testing.T) {
			response := fileUtils.InitUtils()
			if response {
				assert.Equal(t, tt.flag, response)
			} else {
				assert.Equal(t, tt.flag, response)
			}
		})
	}
}

func TestPromptGetYesOrNoInput(t *testing.T) {

	mockPC := fileUtils.PromptContent{"Yes", "msg"}

	cases := []struct {
		name   string
		result bool
		err    bool
	}{
		{
			name:   "Success-true",
			result: true,
			err:    false,
		},
		{
			name:   "Success-false",
			err:    false,
			result: false,
		},
		{
			name:   "Fail",
			err:    true,
			result: false,
		},
	}
	for _, tt := range cases {

		var prompt *promptui.Select
		patchRun := gomonkey.ApplyMethod(
			reflect.TypeOf(prompt),
			"Run",
			func(*promptui.Select) (int, string, error) {
				if !tt.err {
					if tt.result {
						return 1, "Yes", nil
					} else {
						return 1, "No", nil
					}
				} else {
					return 0, "", fmt.Errorf("Error in run")
				}
			},
		)
		defer patchRun.Reset()

		t.Run(tt.name, func(t *testing.T) {
			response := fileUtils.PromptGetYesOrNoInput(mockPC)
			if response {
				assert.Equal(t, true, response)
			} else {
				assert.Equal(t, false, response)
			}
		})
	}
}
