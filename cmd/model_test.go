package cmd_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/wednesday-solutions/negt/cmd"
	"github.com/wednesday-solutions/negt/gqlgenUtils/modelUtils"
)

func TestRunModelE(t *testing.T) {
	cases := []struct {
		name string
		err  bool
	}{
		{
			name: "Success",
			err:  false,
		},
		{
			name: "Fail",
			err:  true,
		},
	}
	for _, tt := range cases {

		patchCreateNewModel := gomonkey.ApplyFunc(
			modelUtils.CreateNewModel,
			func() error {
				if tt.err {
					return fmt.Errorf("Error in CreateNewModel")
				} else {
					return nil
				}
			},
		)
		defer patchCreateNewModel.Reset()

		t.Run(tt.name, func(t *testing.T) {
			err := cmd.RunModelE(&cobra.Command{}, []string{})
			if err != nil {
				assert.Equal(t, true, err != nil)
				assert.Equal(t, true, strings.Contains(err.Error(), "Error in CreateNewModel"))
			} else {
				assert.Equal(t, true, err == nil)
			}
		})
	}
}

func TestModelCmd(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		response := cmd.ModelCmdFn()
		assert.Equal(t, true, response != nil)
	})
}

func TestModelInit(t *testing.T) {

	patchCreateNewModel := gomonkey.ApplyFunc(
		modelUtils.CreateNewModel,
		func() error {
			return nil
		},
	)
	defer patchCreateNewModel.Reset()

	gqlgenCmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {},
	}
	pacthAddCommand := gomonkey.ApplyMethod(
		reflect.TypeOf(gqlgenCmd),
		"AddCommand",
		func(*cobra.Command, ...*cobra.Command) {},
	)
	defer pacthAddCommand.Reset()

	t.Run("Success", func(*testing.T) {
		testing.Init()
	})
}
