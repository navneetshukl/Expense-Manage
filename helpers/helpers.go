package helpers

import "strconv"

//! StringToInt this will convert string to int
func StringToInt(str string) (int, error) {
	if len(str) == 0 {
		return 0, nil
	}
	val, err := strconv.Atoi(str)

	if err != nil {
		return 0, err
	} else {
		return val, nil
	}
}

//! IntToString convert int to string
func IntToString(val int) string {
	str := strconv.Itoa(val)
	return str
}
