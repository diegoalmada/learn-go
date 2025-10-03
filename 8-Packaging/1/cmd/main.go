package main

import (
	"fmt"
	"github.com/diegoalmada/goexpert/8-Packaging/1/math"
)

func main() {
	m := math.NewMath(2, 3)
	fmt.Println(m.Add())
}
