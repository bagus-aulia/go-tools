package general_test

import (
	"testing"

	"github.com/bagus-aulia/go-tools/tools/general"
	"github.com/go-playground/assert/v2"
)

func TestIndexOf(t *testing.T) {
	t.Run("Index string not found", func(t *testing.T) {
		strList := []string{"1.1", "1.2", "1.3"}

		res := general.IndexOf(strList, "1.4")
		assert.Equal(t, -1, res)
	})

	t.Run("Index string found", func(t *testing.T) {
		strList := []string{"1.1", "1.2", "1.3"}

		res := general.IndexOf(strList, "1.2")
		assert.Equal(t, 1, res)
	})

	t.Run("Index int not found", func(t *testing.T) {
		intList := []int{1, 2, 3}

		res := general.IndexOf(intList, 4)
		assert.Equal(t, -1, res)
	})

	t.Run("Index int found", func(t *testing.T) {
		intList := []int{1, 2, 3}

		res := general.IndexOf(intList, 2)
		assert.Equal(t, 1, res)
	})
}
