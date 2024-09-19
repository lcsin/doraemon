package main

import "fmt"

func main() {
	k := "col_0"
	i := 1
	s := fmt.Sprintf("%v#%02d", k, i)
	fmt.Println(s)
}
