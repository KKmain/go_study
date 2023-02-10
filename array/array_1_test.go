package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_array_sort(T *testing.T) {
	fmt.Println(array_sort())
}

func Test_array_sort2(T *testing.T) {
	i := array_sort()
	assert.Equal(T, i, 2, "Test_array_sort2 error")
}
