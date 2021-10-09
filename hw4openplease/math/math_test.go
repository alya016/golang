package math
func TestCaloriesByMeal(t *testing.T){
	a,b := 1,2
	calories := CaloriesByMeal(a, b)
	assert.Equal(t, 222, calories)
}