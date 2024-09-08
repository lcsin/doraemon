package judge

import (
	"regexp"
)

// Number 判断文本是否为纯数字
func Number(text string) bool {
	if len(text) == 0 {
		return false
	}
	matched, _ := regexp.MatchString("^[0-9]*$", text)
	return matched
}

// ContainChinese 判断文本是否包含中文
func ContainChinese(text string) bool {
	if len(text) == 0 {
		return false
	}
	matched, _ := regexp.MatchString("\\p{Han}", text)
	return matched
}

// BalancedPlaceholder 判断文本中的占位符对是否左右匹配
func BalancedPlaceholder(text string, dict map[string]string) bool {
	if len(text) == 0 {
		return false
	}

	stack := make([]string, 0, len(text))
	left := make(map[string]bool, len(dict)/2)
	right := make(map[string]bool, len(dict)/2)
	reverse := make(map[string]string, len(dict))
	for k, v := range dict {
		reverse[v] = k
	}

	for k, v := range dict {
		left[k] = true
		right[v] = true
	}

	for _, v := range text {
		str := string(v)
		switch true {
		case left[str]: // 左边直接入栈
			stack = append(stack, str)
		case right[str]: // 右边判断是否匹配
			if len(stack) == 0 || stack[len(stack)-1] != reverse[str] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
