package main

import (
	"math"
)

type side int

const (
	a = 0
	b = 3
	c = 4
)

func CalcSquare(sideLen float64, sidesNum side) float64 {
	pow := math.Pow(sideLen, 2)
	switch sidesNum {
	case a:
		return math.Pi * pow
	case b:
		return (pow * math.Sqrt(3)) / 4
	case c:
		return pow
	default:
		return 0
	}
}
