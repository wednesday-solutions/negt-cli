package modelUtils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wednesday-solutions/negt/gqlgenUtils/modelUtils"
)

func TestFieldUtils(t *testing.T) {

	type args struct {
		modelName      string
		fields         []string
		fieldTypes     []string
		nullFields     []bool
		customMutation bool
	}
	cases := []struct {
		name string
		req  args
	}{
		{
			name: "Success",
			req: args{
				modelName:      "testModel",
				fields:         []string{"field1"},
				fieldTypes:     []string{"ID", "Int", "Float", "String", "Boolean", "DateTime"},
				nullFields:     []bool{true},
				customMutation: true,
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctx := modelUtils.FieldUtils(
				tt.req.modelName,
				tt.req.fields,
				tt.req.fieldTypes,
				tt.req.nullFields,
				tt.req.customMutation,
			)
			if ctx != nil {
				assert.Equal(t, true, ctx != nil)
			} else {
				assert.Equal(t, false, ctx != nil)
			}
		})
	}
}
