package helper

var MANAGER = "Manager"
var TECHNICIAN = "Technician"

func IsManager(role string) bool {
	return role == MANAGER
}

func IsTechnician(role string) bool {
	return role == TECHNICIAN
}
