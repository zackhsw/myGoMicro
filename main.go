package main

import "fmt"

func main() {
	str := []rune("hello")
	for k, v := range str {
		fmt.Println(k, v, string(v))
	}
	fmt.Println(str)
}
