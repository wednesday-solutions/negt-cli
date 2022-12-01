package cmd_test

import "testing"

func TestInitTest(t *testing.T) {
	t.Run("Success", func(*testing.T) {
		testing.Init()
	})
}
