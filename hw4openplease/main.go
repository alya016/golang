package main

import (
    "fmt"
	"homework/stringt"
	"homework/math"
	"homework/facility"
	"encoding/json"
	"log"
)

type Output struct{
	Isprofitable bool `json:"facility_is_profitable"`
	Profit float64 `json:"profit"`
}


func main(){

	// facility2 := facility.Facility{
	// 	Town: "Merke",
	// 	Coordinates: []float64 {2.00,3.00,8.00,9.00,},
	// 	SalaryToGuard: 8000,
	// 	PaymentFromClient: 8020,
	// }
	facility1 := facility.Facility{
		Town: "Shu",
		Coordinates: []float64 {2.00,3.00,8.00,9.00,},
		SalaryToGuard: 10000.00,
		PaymentFromClient: 12000.00,
	}



//проверяем наличие филиала в списке обслуживаемых на данный момент если необслуживается
// нет смысла проверять рентабильность, выходим из программы
	if stringt.CheckExcistenceFacility(facility1.Town) == false{
		fmt.Printf("%s is not current facility of our company",facility1.Town)
		return }
//отправляем в math рассчитать рентабельность и записываем в переменные
	profitability,profit := math.CheckProfitability(&facility1)

//заполняем структуру и форматируем в json формат, проверяя на ошибку чтобы данные не были пустыми
	output := Output{
		Isprofitable: profitability,
		Profit: profit,
	}
	json_data, err := json.Marshal(output)
	if err != nil {
		log.Fatal(err)
	}

//программа выводит json данные  	
	fmt.Println(string(json_data))
	math.TestCheckProfitability(true,1787.8679656440363)
}

	