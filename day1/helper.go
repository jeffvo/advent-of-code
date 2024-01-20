package main

func IsDigit(s byte) bool {
	if s >= 48 && s <= 57 {
		return true
	}

	return false
}

func ToInt(c byte) int {
	return int(c - '0')
}
