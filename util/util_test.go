package util

import (
	"fmt"
	"testing"
)

func TestUtil(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	slice = DelSliceElement(slice, 3)
	fmt.Println(slice)
}

func TestRemoveDuplicates(t *testing.T) {
	slice := []int{1, 2, 3, 4, 4, 5, 6}
	slice = RemoveDuplicates(slice)
	fmt.Println(slice)
}
