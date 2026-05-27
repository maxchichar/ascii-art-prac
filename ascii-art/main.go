package main

import(
	"fmt"
	"os"
	"strings"
)


func main()  {
	if len(os.Args) < 3 {
		fmt.Println("usage: go run . [STRING] [BANNER]")
	}
	
	bannername := os.Args[2]

	var banner = bannername + ".txt"
	
	
	switch bannername {
	case "thinkertoy":
		data, err:= os.ReadFile(banner)
		if err != nil {
			fmt.Println("error:", err)
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

	case "shadow":
		data, err := os.ReadFile(banner)
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
	case "standard":
		data, err := os.ReadFile(banner)
		if err != nil{
			fmt.Println("error: ", err)
			return
		}
		
		
		content := string(data)
		chunk := strings.ReplaceAll(content, "\"\n", "\n")
		lines := strings.Split(chunk, "\n")
	
		for row := 0; row < 8; row++ {
			for _, ch := range os.Args[1]{
				ascii := (ch - 32) * 9 + 1
				fmt.Print(lines[int(ascii) + row])
			}
			fmt.Println()
		}
	default:
		fmt.Println("usage: go run . [STRING] [BANNER]")
	}
	
	
}

// func Processor()