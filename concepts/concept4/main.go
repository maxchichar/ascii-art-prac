/*
🟡 Exercise 4 — Load the banner file
Task:

Download the standard.txt banner file from the ascii-art project
Read it using os.ReadFile
Split it into lines using strings.Split
Using your formula, extract the 8 lines for one hardcoded character of your choice
Print those 8 lines to confirm you got the right art

Things to figure out yourself:

What package do you need for ReadFile?
What package do you need for Split?
Once you have lines as a slice, how do you extract exactly 8 lines starting at your formula's index?

You already know how to slice a slice. You did it in Exercise 2.

One rule — before printing, also print the index number your formula gives you for your chosen character, so you can verify the math is right.
Go.
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("shadow.txt")
	if err != nil{
		fmt.Println("Error:", err)
		return
	}

	content := string(file)
	lines := strings.Split(content, "\n")

	for row := 0; row < 8; row++ {
		for _, ch := range "Hello Ascii-Art"{
			start := (ch - 32) * 9 + 1
			fmt.Print(lines[int(start) + row])
		}
		fmt.Println()
	}
}

