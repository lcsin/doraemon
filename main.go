package main

import (
	"fmt"
)

var (
	reg1 = `\[\[.*?\]\]|\[.*?\]`
	reg2 = `\{\{.*?\}\}|\{.*?\}`
)

func main() {
	//text := "这是一段包含[[姓名]]的[文本]。"
	text := "这是一段包含((姓名))的(文本)。这里有{花括号}和<尖括号>"

	result := extractTextBetweenWildcards(text, "(", ")")

	fmt.Println("提取的文本：")
	for _, match := range result {
		fmt.Println(match)
	}
}

func extractTextBetweenWildcards(text, start, end string) []string {
	var result []string
	var stack []int

	for i := 0; i < len(text); i++ {
		if i+len(start) <= len(text) && text[i:i+len(start)] == start {
			stack = append(stack, i)
			i += len(start) - 1
		} else if i+len(end) <= len(text) && text[i:i+len(end)] == end && len(stack) > 0 {
			startIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			extracted := text[startIndex : i+len(end)]
			result = append(result, extracted)
			i += len(end) - 1
		}
	}

	return result
}
