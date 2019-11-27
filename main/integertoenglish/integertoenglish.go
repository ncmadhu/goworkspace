package integertoenglish

//NumberToWords - convert number to words
func NumberToWords(num int) string {
	if num < 1 {
		return "Zero"
	}
	return getBillions(num)

}

func getBillions(num int) string {

	quotient := num / 1000000000
	remainder := num % 1000000000
	billions := ""
	if quotient > 0 {
		billions = getHundreds(quotient) + " Billion"
	}
	if remainder > 0 {
		millions := getMillions(remainder)
		if billions != "" && millions != "" {
			billions = billions + " " + millions
		} else if millions != "" {
			billions = millions
		}
	}
	return billions
}

func getMillions(num int) string {

	quotient := num / 1000000
	remainder := num % 1000000
	millions := ""
	if quotient > 0 {
		millions = getHundreds(quotient) + " Million"
	}
	if remainder > 0 {
		thousands := getThousands(remainder)
		if millions != "" && thousands != "" {
			millions = millions + " " + thousands
		} else if thousands != "" {
			millions = thousands
		}
	}
	return millions
}

func getThousands(num int) string {

	quotient := num / 1000
	remainder := num % 1000
	thousands := ""
	if quotient > 0 {
		thousands = getHundreds(quotient) + " Thousand"
	}
	if remainder > 0 {
		hundreds := getHundreds(remainder)
		if thousands != "" && hundreds != "" {
			thousands = thousands + " " + hundreds
		} else if hundreds != "" {
			thousands = hundreds
		}
	}
	return thousands
}

func getHundreds(num int) string {

	quotient := num / 100
	remainder := num % 100
	hundreds := ""
	if quotient > 0 {
		hundreds = getOnes(quotient) + " Hundred"
	}
	if remainder > 0 {
		tens := getTens(remainder)
		if hundreds != "" && tens != "" {
			hundreds = hundreds + " " + tens
		} else if tens != "" {
			hundreds = tens
		}
	}
	return hundreds
}

func getTens(num int) string {
	tenMap := map[int]string{
		10: "Ten",
		11: "Eleven",
		12: "Twelve",
		13: "Thirteen",
		14: "Fourteen",
		15: "Fifteen",
		16: "Sixteen",
		17: "Seventeen",
		18: "Eighteen",
		19: "Nineteen",
		2:  "Twenty",
		3:  "Thirty",
		4:  "Forty",
		5:  "Fifty",
		6:  "Sixty",
		7:  "Seventy",
		8:  "Eighty",
		9:  "Ninety",
	}
	tens := ""
	if num < 10 {
		return getOnes(num)

	}
	if val, ok := tenMap[num]; ok {
		tens = val
	} else {
		quotient := num / 10
		if quotient > 0 {
			tens = tenMap[quotient]
		}
		remainder := num % 10
		if remainder > 0 {
			ones := getOnes(remainder)
			if tens != "" {
				tens = tens + " " + ones
			} else {
				tens = ones
			}
		}
	}
	return tens
}

func getOnes(num int) string {
	numMap := map[int]string{
		0: "Zero",
		1: "One",
		2: "Two",
		3: "Three",
		4: "Four",
		5: "Five",
		6: "Six",
		7: "Seven",
		8: "Eight",
		9: "Nine",
	}
	return numMap[num]
}
