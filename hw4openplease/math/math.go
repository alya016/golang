package math

import (
	"math"
	"homework/facility"
)



func CheckProfitability(f *facility.Facility) bool{
	distance := CalculateDistance(f.Coordinates)
	const k = 25.00 //коэфицент затраты бензина на расстояние
	difference := f.PaymentFromClient - (distance*k + f.SalaryToGuard)
	if difference >= 1000{
		return true
	} else {return false}
}
func CalculateDistance(x []float64)float64{
	Square1 := 0.00
	Square2 := 0.00
	for i := 0; i<1; i++{
		Square1 = math.Pow(math.Abs(x[i] - x[i+2]), 2)
		Square2 = math.Pow(math.Abs(x[i+1] - x[i+3]), 2)
	}
	return math.Sqrt(Square1+Square2)
}