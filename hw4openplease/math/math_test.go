package math

import (
	"homework/facility"
	"testing"
)

func TestCheckProfitability(t *testing.T){
	testCase := TestCase{
		expectedProfitability: true,
		expectedProfit: 1787.8679656440363,
	}
	facility1 := facility.Facility{
		Town: "Shu",
		Coordinates: []float64 {2.00,3.00,8.00,9.00,},
		SalaryToGuard: 10000.00,
		PaymentFromClient: 12000.00,
	}
	actualProfitability,actualProfit := CheckProfitability(&facility1)
	if actualProfitability != testCase.expectedProfitability ||  actualProfit != testCase.expectedProfit{
		t.Fail()
	}
}