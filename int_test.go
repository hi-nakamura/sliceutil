package sliceutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicateInts(t *testing.T) {
	data := []int{1, 1, 2, 3, 2, 1, 1, 1, 1, 1, 2, 2, 3, 4}

	as := assert.New(t)
	result := removeDuplicateInts(data)
	as.Len(result, 4)
	as.Equal([]int{1, 2, 3, 4}, result)
}
