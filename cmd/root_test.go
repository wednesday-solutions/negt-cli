package cmd_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/wednesday-solutions/negt/cmd"
)

func TestRunTestCmd(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		cmd.RunTestCmd(&cobra.Command{}, []string{}) // nolint:errcheck
	})
}

func TestExecute(t *testing.T) {

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

		var rootCmd *cobra.Command
		patchExecute := gomonkey.ApplyMethod(
			reflect.TypeOf(rootCmd),
			"Execute",
			func(*cobra.Command) error {
				if tt.err {
					return fmt.Errorf("Error in execute")
				} else {
					return nil
				}
			},
		)
		defer patchExecute.Reset()

		t.Run(tt.name, func(t *testing.T) {
			err := cmd.Execute()
			if tt.err {
				assert.Equal(t, true, err != nil)
			} else {
				assert.Equal(t, true, err == nil)
			}
		})
	}
}
