package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

// var c, python, java bool

// Variables with initializers
//var i, j int = 1, 2

// Basic types outer
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 13))
	//a, b := swap("hello", "world")
	//fmt.Println(a, b)
	fmt.Println(split(17))
	// var i int
	// fmt.Println(i, c, python, java)
	// var c, python, java = true, false, "no!"
	// fmt.Println(i, j, c, python, java)

	//Short variable declarations
	//var i, j int = 1, 2
	//k := 3
	//c, python, java := true, false, "no!"
	//fmt.Println(i, j, k, c, python, java)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// // Zero values
	// var i int
	// //var f float64
	// var b bool
	// var s string
	// //fmt.Printf("%v %v %v %q\n", i, f, b, s)

	//Type conversions
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	// Type inference
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)

	// Constants
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
