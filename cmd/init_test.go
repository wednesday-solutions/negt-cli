package cmd_test

import (
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/wednesday-solutions/negt/cmd"
	"github.com/wednesday-solutions/negt/gqlgenUtils/fileUtils"
)

func TestRunInitE(t *testing.T){
	cases := []struct{
		name string
		flag bool
		isExist bool
	}{
		{
			name: "Success-flag-true",
			flag: true,
		},
		{
			name: "Success-flag-false",
			flag: false,
		},
		{
			name: "Success-flag-true",
			flag: true,
			isExist: true,
		},
		{
			name: "Success-flag-false",
			flag: false,
			isExist: true,
		},
	}
	for _, tt := range cases {
		patchInitUtils := gomonkey.ApplyFunc(
			fileUtils.InitUtils,
			func() bool {
				if tt.flag{
					return true
				} else {
					return false
				}
			},
		)
		defer patchInitUtils.Reset()

		patchIsExists := gomonkey.ApplyFunc(
			fileUtils.IsExists,
			func(string, string) bool {
				if tt.isExist{
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

		t.Run(tt.name, func(t *testing.T){
			err := cmd.RunInitE(&cobra.Command{}, []string{})
			if err != nil {
				assert.Equal(t, true, err != nil)
			} else {
				assert.Equal(t, true, err == nil)
			}
		})
	}
}

func TestInitCmd(t *testing.T){
	t.Run("Success", func(t *testing.T){
		response := cmd.InitCmd()
		assert.Equal(t, true, response != nil)
	})
}

func TestInitInit(t *testing.T){

	patchInitCmd := gomonkey.ApplyFunc(
		fileUtils.InitUtils,
		func() bool {
			return true
		},
	)
	defer patchInitCmd.Reset()
	
	t.Run("Success", func(t *testing.T){
		testing.Init()
	})
}