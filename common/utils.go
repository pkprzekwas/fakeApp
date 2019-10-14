package common

func BoolToInt(a bool) int {
	if a {
		return 1
	}
	return 0
}

func IntToBool(a int) bool {
	if a != 0 {
		return true
	}
	return false
}
