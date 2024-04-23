package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	// Structs
	fmt.Println(Vertex{1, 2})
}
