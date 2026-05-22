/*	
	🟢 Exercise 2
	Write a program that:

	Has a hardcoded slice of 8 strings (just make up any 8 strings)
	Prints only lines at index 1, 2, 3 using a single slice operation
	Then loops over the full slice and prints each line with its index

	Expected output format:
	-- Slice cut --
	line3
	line4
	line5

	-- Full loop --
	0 line1
	1 line2
	...
*/	

package main

import(
	"fmt"
)

func main()  {
	lines := []string{"line1", "line2", "line3", "line4", "line5", "line6", "line7", "line8"}
	
	fmt.Println(lines[1:4])

	for i, line := range lines{
		fmt.Println(i, string(line))
	}
}