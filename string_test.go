package sliceutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicateStrings(t *testing.T) {
	data := []string{"test1", "test1", "test2", "test3", "test2", "test1", "test1", "test1", "test1", "test1", "test2", "test2", "test3", "test4"}

	as := assert.New(t)
	result := removeDuplicateStrings(data)
	as.Len(result, 4)
	as.Equal([]string{"test1", "test2", "test3", "test4"}, result)
}
