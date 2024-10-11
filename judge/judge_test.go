package judge

import (
	"fmt"
	"testing"
)

func TestBalancedHolder(t *testing.T) {
	dict := make(map[string]string, 0)
	dict["{"] = "}"
	dict["["] = "]"
	dict["<"] = ">"

	text := "a{nihao}b,c[nihao]d,<nihao>.aaa"
	fmt.Println(BalancedPlaceholder(text, dict)) // true

	text = "a{nihao},b{nihao},c[<sdfsf>]"
	fmt.Println(BalancedPlaceholder(text, dict)) // true

	text = "{<>[]<}}"
	fmt.Println(BalancedPlaceholder(text, dict)) // false
}
