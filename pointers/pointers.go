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

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// pointers to structs
	point := &v
	point.X = 1e9
	fmt.Println(v)
}
