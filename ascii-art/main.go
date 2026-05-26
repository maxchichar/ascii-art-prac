package main

import(
	"fmt"
	"os"
	"strings"
)


func main()  {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run . [STRING] [BANNER]")
	}
	
	data, err := os.ReadFile("standard.txt")
	if err != nil{
		fmt.Println("error: ", err)
		return
	}

	content := string(data)
	lines := strings.Split(content, "\n")

	for row := 0; row < 8; row++ {
		for _, ch := range os.Args[1]{
			ascii := (ch - 32) * 9 + 1
			fmt.Print(lines[int(ascii) + row])
		}
		fmt.Println()
	}
}