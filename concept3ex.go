/*
		🟡 Exercise 3 — Spicy
	You have this data:

	'H' → [" _ _ ", "| | | |", "| |__| |", "|  __  |", "| |  | |", "|_|  |_|", "      ", "      "]
	'i' → ["  ", " _ ", "(_)", "| |", "| |", "|_|", "   ", "   "]

	Task:

	Store both characters in a map where the key is a rune
	Take the hardcoded string "Hi"
	Loop over the string
	For each of the 8 rows, print both characters side by side on the same line

	The twist — think before you code: how do you print row 0 of H and row 0 of i on the same line, then row 1 of H and row 1 of i, and so on? You need two loops. Figure out which one is outer and which is inner and why.
	Go.
*/

package main

import(
	"fmt"
)

func main()  {
	char := map[rune][]string{
		'H': []string{" _ _ ", "| | | |", "| |__| |", "|  __  |", "| |  | |", "|_|  |_|", "      ", "      "},
		'i': []string{"  ", " _ ", "(_)", "| |", "| |", "|_|", "   ", "   "},
	}

	// fmt.Print(char['H'], char['i'])

	for row := 0; row <= 7; row++{
		// fmt.Print(char['H'], char['i'])
		for _, ch := range char['H']{
			fmt.Print(ch)
		}
	}
	// fmt.Print(len(char['H']))
}