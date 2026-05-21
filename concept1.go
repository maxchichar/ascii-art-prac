package main

import(
	"fmt"
)

func main()  {
	word := "Hello"

	for i, ch := range word{
		fmt.Println(i, string(ch))
	}
}
