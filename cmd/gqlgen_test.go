package cmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wednesday-solutions/negt/cmd"
)

func TestInitGqlgen(t *testing.T) {
	t.Run("Success", func(*testing.T) {
		testing.Init()
	})
}

func TestGqlgenCmd(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		response := cmd.GqlgenCmdFn()
		assert.Equal(t, true, response != nil)
	})
}
