package Radian

import (
	"math"
)

func DegreeToRadian(degree float64) float64 {
	return degree * (math.Pi / 180)
}

func RadianToDegree(radian float64) float64 {
	return radian * (180 / math.Pi)
}
