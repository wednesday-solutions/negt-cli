package cmd_test

import "testing"

func TestInitTest(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		testing.Init()
	})
}
