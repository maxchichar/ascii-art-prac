package main

import(
	"fmt"
)

func main()  {
	m := map[string]string{
		"a": "alpha",
		"b": "beta",
	}
	
	c := map[string]string{
		"z": "zebra",
		"a": "altraford",
		"b": "benita",
	}

	fmt.Println(m["b"])

	value, ok := c["j"]

	if ok {
		fmt.Println("Found:", value)
	}else{
		fmt.Println("Value not available")
	}

	g := map[string][]string{
		"HI": []string{"line1", "line2", "line3"},
	}

	fmt.Println(g["HI"])

}
