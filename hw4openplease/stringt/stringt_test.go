package stringt

import(
	"homework/facility"
	"testing"
)

func TestCheckExcistenceFacility(t *testing.T){
	facility1 := facility.Facility{
		Town: "Shu",
		Coordinates: []float64 {2.00,3.00,8.00,9.00,},
		SalaryToGuard: 10000.00,
		PaymentFromClient: 12000.00,
	}
	expectedIsExist := true
	actualIsExist := CheckExcistenceFacility(facility1.Town)
	if expectedIsExist != actualIsExist{
		t.Fail()
	}
}