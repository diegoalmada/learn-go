package main

import (
	"github.com/diegoalmada/learn-go/8-Packaging/4/math"
	"github.com/google/uuid"
)

func main() {

	m := math.NewMath(10, 2)
	println(m.Add())
	println(uuid.New().String())
}
