package leap

const testVersion = 3

func IsLeapYear(year int) bool {
	// Write some code here to pass the test suite.

	var status bool

	if year % 4 == 0 {
		if year % 100 != 0 {
			status = true
			return status
		}
		if year % 100 == 0 && year % 400 != 0 {
			status = false
			return status
		} else if year % 100 == 0 && year % 400 == 0 {
			status = true
			return status
		}
	} else {
		status = false
		return status
	}
 return status
}