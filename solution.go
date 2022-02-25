package square

import (
	"math"
)

type side int

const (
	SidesCircle   = 0
	SidesTriangle = 3
	SidesSquare   = 4
)

func CalcSquare(sideLen float64, sidesNum side) float64 {
	pow := math.Pow(sideLen, 2)
	switch sidesNum {
	case SidesCircle:
		return math.Pi * pow
	case SidesTriangle:
		return (pow * math.Sqrt(3)) / 4
	case SidesSquare:
		return pow
	default:
		return 0
	}
}
