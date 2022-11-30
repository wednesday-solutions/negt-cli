package hbs_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/wednesday-solutions/negt/gqlgenUtils/fileUtils"
	"github.com/wednesday-solutions/negt/hbs"
)

func TestMutationSource(t *testing.T) {
	type args struct {
		modelName, path, file string
		ctx                   map[string]interface{}
	}
	cases := []struct {
		name      string
		req       args
		err       bool
		genTplErr bool
	}{
		{
			name: "Success",
			req: args{
				modelName: "modelName",
				path:      "path",
				file:      "file",
				ctx:       map[string]interface{}{},
			},
			err: false,
		},
		{
			name: "Fail",
			req: args{
				modelName: "modelName",
				path:      "path",
				file:      "file",
				ctx:       map[string]interface{}{},
			},
			err:       true,
			genTplErr: true,
		},
	}

	for _, tt := range cases {
		patchGenTpl := gomonkey.ApplyFunc(
			hbs.GenerateTemplate,
			func(string, map[string]interface{}) (string, error) {
				if tt.genTplErr {
					return "", fmt.Errorf("Error in gentpl")
				}
				return "source", nil
			},
		)
		defer patchGenTpl.Reset()

		patchWriteToFile := gomonkey.ApplyFunc(
			fileUtils.WriteToFile,
			func(string, string, string) error {
				return nil
			},
		)
		defer patchWriteToFile.Reset()

		t.Run(tt.name, func(t *testing.T) {
			err := hbs.MutationSource(
				tt.req.modelName,
				tt.req.path,
				tt.req.file,
				tt.req.ctx,
			)
			if err != nil {
				assert.Equal(t, true, err != nil)
				if tt.genTplErr {
					assert.Equal(t, true, strings.Contains(err.Error(), "Error in gentpl"))
				}
			} else {
				assert.Equal(t, true, err == nil)
			}
		})
	}
}
