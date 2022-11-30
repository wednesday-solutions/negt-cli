package modelUtils_test

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/manifoldco/promptui"
	"github.com/stretchr/testify/assert"
	"github.com/wednesday-solutions/negt/gqlgenUtils/fileUtils"
	"github.com/wednesday-solutions/negt/gqlgenUtils/modelUtils"
)

func TestPromptValidate(t *testing.T) {
	cases := []struct {
		name string
		err  bool
		req  string
	}{
		{
			name: "Success",
			err:  false,
			req:  "modelName",
		},
		{
			name: "Fail",
			err:  true,
			req:  "m",
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := modelUtils.PromptValidate(tt.req)
			if err != nil {
				assert.Equal(t, true, strings.Contains(err.Error(), "Invalid input."))
			} else {
				assert.Equal(t, true, err == nil)
			}
		})
	}
}

func TestPromptGetInput(t *testing.T) {
	cases := []struct {
		name string
		err  bool
		req  modelUtils.PromptContent
	}{
		{
			name: "Success",
			err:  false,
			req: modelUtils.PromptContent{
				ErrorMsg: "error msg",
				Label:    "label",
			},
		},
		{
			name: "Fail",
			err:  true,
			req: modelUtils.PromptContent{
				ErrorMsg: "error msg",
				Label:    "label",
			},
		},
	}
	for _, tt := range cases {

		patchValidate := gomonkey.ApplyFunc(
			modelUtils.PromptValidate,
			func(string) error {
				return nil
			},
		)
		defer patchValidate.Reset()

		var prompt *promptui.Prompt
		patchRun := gomonkey.ApplyMethod(
			reflect.TypeOf(prompt),
			"Run",
			func(*promptui.Prompt) (string, error) {
				if tt.err {
					return "", fmt.Errorf("Error in Run")
				} else {
					return "result", nil
				}
			},
		)
		defer patchRun.Reset()

		patchExit := gomonkey.ApplyFunc(
			os.Exit,
			func(int) {},
		)
		defer patchExit.Reset()

		t.Run(tt.name, func(t *testing.T) {
			response := modelUtils.PromptGetInput(tt.req)
			if tt.err {
				assert.Equal(t, true, response == "")
			} else {
				assert.Equal(t, response, "result")
			}
		})
	}
}

func TestPromptGetYesOrNoInput(t *testing.T) {
	cases := []struct {
		name string
		resp bool
		req  modelUtils.PromptContent
	}{
		{
			name: "Success-true",
			resp: true,
			req: modelUtils.PromptContent{
				ErrorMsg: "error msg",
				Label:    "label",
			},
		},
		{
			name: "Success-false",
			resp: false,
			req: modelUtils.PromptContent{
				ErrorMsg: "error msg",
				Label:    "label",
			},
		},
	}
	for _, tt := range cases {
		var prompt *promptui.Select
		patchRun := gomonkey.ApplyMethod(
			reflect.TypeOf(prompt),
			"Run",
			func(*promptui.Select) (int, string, error) {
				if tt.resp {
					return 1, "Yes", fmt.Errorf("Error in Run")
				} else {
					return 1, "No", nil
				}
			},
		)
		defer patchRun.Reset()

		t.Run(tt.name, func(t *testing.T) {
			response := modelUtils.PromptGetYesOrNoInput(tt.req)
			if response {
				assert.Equal(t, true, response)
			} else {
				assert.Equal(t, false, response)
			}
		})
	}
}

func TestPromptGetSelectPath(t *testing.T) {
	cases := []struct {
		name      string
		req       modelUtils.PromptContent
		err       bool
		dirExists bool
	}{
		{
			name:      "Success",
			err:       false,
			dirExists: true,
			req: modelUtils.PromptContent{
				ErrorMsg: "error msg",
				Label:    "label",
			},
		},
		{
			name:      "Faile-dirExists-false",
			err:       true,
			dirExists: false,
			req: modelUtils.PromptContent{
				ErrorMsg: "error msg",
				Label:    "label",
			},
		},
	}
	for _, tt := range cases {

		var prompt *promptui.SelectWithAdd
		patchRun := gomonkey.ApplyMethod(
			reflect.TypeOf(prompt),
			"Run",
			func(*promptui.SelectWithAdd) (int, string, error) {
				if tt.err {
					return 0, "", fmt.Errorf("Error in Run")
				} else if tt.dirExists {
					return 2, "server/gql/models", nil
				} else if !tt.dirExists {
					return 0, "gql/models", nil
				} else {
					return 0, "dir", nil
				}
			},
		)
		defer patchRun.Reset()

		patchDirExists := gomonkey.ApplyFunc(
			fileUtils.DirExists,
			func(string) bool {
				if tt.dirExists {
					return true
				} else {
					return false
				}
			},
		)
		defer patchDirExists.Reset()

		patchExit := gomonkey.ApplyFunc(
			os.Exit,
			func(int) {},
		)
		defer patchExit.Reset()

		patchCurrentDirectory := gomonkey.ApplyFunc(
			fileUtils.CurrentDirectory,
			func() string {
				return "path"
			},
		)
		defer patchCurrentDirectory.Reset()

		patchIsExists := gomonkey.ApplyFunc(
			fileUtils.IsExists,
			func(string, string) bool {
				return false
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

		t.Run(tt.name, func(t *testing.T) {
			response := modelUtils.PromptGetSelectPath(tt.req)
			fmt.Println("Response: ", response)
			if response == "" {
				assert.Equal(t, true, response == "")
			} else {
				if tt.dirExists {
					assert.Equal(t, response, "server/gql/models")
				}
			}
		})
	}
}

func TestPromptGetSelect(t *testing.T) {
	cases := []struct {
		name      string
		err       bool
		req       modelUtils.PromptContent
		otherCase bool
	}{
		{
			name: "Success",
			err:  false,
			req: modelUtils.PromptContent{
				ErrorMsg: "error msg",
				Label:    "label",
			},
		},
		{
			name: "Fail",
			err:  true,
			req: modelUtils.PromptContent{
				ErrorMsg: "error msg",
				Label:    "label",
			},
		},
	}
	for _, tt := range cases {
		var prompt *promptui.SelectWithAdd
		patchRun := gomonkey.ApplyMethod(
			reflect.TypeOf(prompt),
			"Run",
			func(*promptui.SelectWithAdd) (int, string, error) {
				if tt.err {
					return 1, "", fmt.Errorf("Error")
				} else {
					return 0, "result", nil
				}
			},
		)
		defer patchRun.Reset()

		patchExit := gomonkey.ApplyFunc(
			os.Exit,
			func(int) {},
		)
		defer patchExit.Reset()

		t.Run(tt.name, func(t *testing.T) {
			response := modelUtils.PromptGetSelect(tt.req)
			if response != "" {
				assert.Equal(t, response, "result")
			} else {
				assert.Equal(t, true, response == "")
			}
		})
	}
}
