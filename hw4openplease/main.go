package main

import (
    "fmt"
	"hw4/math"
)

func main(){
	a int := 5
    calories := math.CaloriesByMeal(a)
	fmt.Println(calories)
}