package main

import (
    "fmt"
	"homework/stringt"
	"homework/math"
	"homework/facility"
	// "encoding/json"
)




func main(){

	facility1 := facility.Facility{
		Town: "Shu",
		Coordinates: []float64 {2.00,3.00,8.00,9.00,},
		SalaryToGuard: 10000.00,
		PaymentFromClient: 12000.00,
	}


	if stringt.CheckExcistenceFacility(facility1.Town){
		profitability := math.CheckProfitability(&facility1)
		fmt.Println("Facility is profitable:",profitability)
	} 
	// else {
	// 	fmt.Printf("%s is not current facility of our company",facility1.Town)
	// }

	// profitability, err := json.Marshal(profitability)


}


	// profitability := math.CheckProfitability(facility2)
	// facility2 := Facility{
	// 	Town: "Merke",
	// 	DistancesToGuard: 100,
	// 	SalaryToGuard: 8000,
	// 	PaymentFromClient: 10000,
	// }		