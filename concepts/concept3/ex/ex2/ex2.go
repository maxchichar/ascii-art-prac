/*
## 🟡 Exercise — Spicy

Same concept as before but with a trap inside.

Here's your map data:

- `'G'` → `["  ___  ", " / _ \ ", "| |  ___", "| | |_  |", "| |__| |", " \____/", "       ", "       "]`
- `'o'` → `["      ", "  __  ", " /  \ ", "| () |", " \__/ ", "      ", "      ", "      "]`
- `'!'` → `[" _ ", "| |", "| |", "| |", "|_|", "(_)", "   ", "   "]`

**Task:**
- Store all three in a map
- Your input string is `"Go!"`
- Print it in ASCII art row by row

**The trap:**
Somewhere in your code, add a check — if a character from the string is **not found** in the map, skip it silently instead of crashing.

**Think before coding:**
- What happens if you try to access `char[ch]` and `ch` doesn't exist in the map?
- What tool do you already know that checks if a key exists?

Go.
*/

package main

import(
	"fmt"
)

func main()  {
	char := map[rune][]string{
		'G' : {"  ___  ", ` / _ \ `, "| |  __", "| | |_ |", "| |__| |", ` \____/`, "       ", "       "},
		'o' : {"      ", "   __  ", `  /  \ `, "| () |", ` \__/ `, "      ", "      ", "      "},
		'!' : {"  _ ", "| |", "| |", "| |", "|_|", " (_)", "   ", "   "},
	}

	for row := 0; row < 8; row++{
		for _, ch := range "Go!"{
			art, ok := char[ch]
			if ok{
				fmt.Print(art[row])
			}else{
				fmt.Println(art,"not available")
				return
			}
		}
		fmt.Println()
	}
}
