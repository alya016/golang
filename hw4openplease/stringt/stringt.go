package stringt

import "strings"

func CheckExcistenceFacility(town string) bool{
	CurrentFacilities := "Shu, Almaty, Kordai, Kokkainar, Aktau, Nursultan"
	return strings.Contains(CurrentFacilities,town)
}