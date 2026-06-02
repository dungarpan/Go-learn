package main

import "fmt"

var x int = 42
var y float64 = 30.2
var sum1 float64 = float64(x) + y
var sum2 int = x + int(y)

func main() {
	fmt.Println(sum1, sum2)
	fmt.Println(sum3, sum4)
}
