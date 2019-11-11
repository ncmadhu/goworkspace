package romantointeger

//RomanToInt - convert roman to integer
func RomanToInt(s string) int {
	length := len(s)
	index := 0
	number := 0

	romanMap := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}

	for index < length-1 {
		roman := string(s[index])
		value := romanMap[roman]
		index = index + 1
		nextRoman := roman + string(s[index])
		if temp, ok := romanMap[nextRoman]; ok {
			value = temp
			index = index + 1
		}
		number = number + value
	}
	if index < length {
		number = number + romanMap[string(s[index])]
	}
	return number

}
