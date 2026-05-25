/*
🟢 Exercise — Simple

Remove your hardcoded string
Read it from os.Args[1] instead
Run your program like this:

go run . "Hello"
go run . "Hi There"
That's it. One change. Go.
*/

package main

import(
	"fmt"
	"os"
	"strings"
)

func main()  {
	if len(os.Args) < 2{
		fmt.Println("USAGE: go run . [STRING]")
		return
	}

	file, err := os.ReadFile("standard.txt")
	if err != nil{
		fmt.Println("error: ", err)
		return
	}

	content := string(file)
	lines := strings.Split(content, "\n")

	for row := 0; row < 8; row++ {
		for _, ch := range os.Args[1]{
			start := (ch - 32) * 9 + 1
			fmt.Print(lines[int(start) + row])
		}
		fmt.Println()
	}
	
}

