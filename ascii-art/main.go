package main

import(
	"fmt"
	"os"
	"strings"
)


var availableBannerMessage = `
Available Banners: "standard", "shadow", "thinkertoy"

Usage: 
go run . "Hello" "standard"
go run . "Hello" "shadow"
go run . "Hello" "thinkertoy"
`


func main()  {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("usage: go run . [STRING] [BANNER]")
		return
	}

	var input = os.Args[1]

	var bannername = "standard"

	if len(os.Args) == 3 {
		bannername = os.Args[2]
	}

	lines, err := LoadBanner(bannername + ".txt")
	if err != nil {
		fmt.Print("error:", err)
		fmt.Println(availableBannerMessage)
		return
	}

	if input == "" {
		return
	}
	
	// userInput := strings.ReplaceAll(input)

	for row := 0; row < 8; row++ {
		for _, char := range input{
			ascii_char := (char - 32) * 9 + 1
			fmt.Print(lines[int(ascii_char) + row])
		}
		fmt.Println()
	}
}


func LoadBanner(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	
	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"),"\n")
	
	return lines, nil
}