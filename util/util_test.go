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

func TestExtractTextBetweenWildcards(t *testing.T) {
	text := "这是一段包含[[姓名]]的[文本]"
	fmt.Println(ExtractTextBetweenWildcards(text, "[", "]")) // [[姓名] [[姓名]] [文本]]
}
