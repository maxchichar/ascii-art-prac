package main

import(
	"fmt"
)

func main()  {
	// for loops and strings
	word := "Hello"

	// index guve the exact number of the character
	for i, ch := range word {
		fmt.Println(i, string(ch))
	}


	lines := []string{"line1", "line2", "line3", "line4", "line5", "line6", "line7", "line8"}

	fmt.Println(lines[2:5])

	for i, line := range lines{
		fmt.Println(i, line)
	}
}