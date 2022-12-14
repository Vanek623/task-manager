package service

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsTitleOk(t *testing.T) {
	t.Parallel()
	t.Run("normal", func(t *testing.T) {
		str := strings.Repeat("a", maxNameLen-1)
		assert.NoError(t, isTitleOk(str))
	})

	t.Run("large", func(t *testing.T) {
		str := strings.Repeat("a", maxNameLen+1)
		assert.Error(t, isTitleOk(str))
	})

	t.Run("empty", func(t *testing.T) {
		assert.Error(t, isTitleOk(""))
	})
}

func TestDescriptionOk(t *testing.T) {
	t.Parallel()
	t.Run("normal", func(t *testing.T) {
		str := strings.Repeat("a", maxDescriptionLen-1)
		assert.NoError(t, isDescriptionOk(str))
	})

	t.Run("large", func(t *testing.T) {
		str := strings.Repeat("a", maxDescriptionLen+1)
		assert.Error(t, isDescriptionOk(str))
	})

	t.Run("empty", func(t *testing.T) {
		assert.NoError(t, isDescriptionOk(""))
	})
}
