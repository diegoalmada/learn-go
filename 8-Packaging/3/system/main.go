package main

import "github.com/diegoalmada/learn-go/8-Packaging/3/math"

func main() {

	m := math.NewMath(10, 2)
	println(m.Add())
}
