package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/wednesday-solutions/negt/cmd"
)


func TestMain(t *testing.T) {
	cases := []struct{
		name string
		err bool
	}{
		{
			name: "Success",
			err: false,
		},
		{
			name: "Fail",
			err: true,
		},
	}
	for _, tt := range cases {
		patchExit := gomonkey.ApplyFunc(
			os.Exit,
			func(int){},
		)
		defer patchExit.Reset()

		patchExecute := gomonkey.ApplyFunc(
			cmd.Execute,
			func() error {
				if tt.err {
					return fmt.Errorf("Error in Execute")
				} else {
					return nil
				}
			},
		)
		defer patchExecute.Reset()

		t.Run("Success", func(t *testing.T){
			main()
		})
	}
}
