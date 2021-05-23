package util

func If(condition bool, value1, value2 interface{}) interface{} {
	if condition {
		return value1
	}
	return value2
}
