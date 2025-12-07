package types

import "math"

const numPower = 2 // он типо untyped так как может в float перейти

func get2Abs(value int) int {
	result := math.Pow(float64(value), numPower)
	return int(math.Abs(result))
}
