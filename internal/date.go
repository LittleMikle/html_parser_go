package internal

import (
	"fmt"
	"time"
)

var MaxDateUSD string
var MaxDateEUR string

func GenerateDate() {
	x := []rune("13.04.2023")
	for i := 0; i < 3; i++ {
		time.Sleep(10 * time.Second)
		valUSD, valEUR := DataParser(string(x))
		if valUSD > USD {
			USD = valUSD
			MaxDateUSD = string(x)
		}
		avgUSD += valUSD

		if valEUR > EUR {
			EUR = valEUR
			MaxDateEUR = string(x)
		}
		avgEUR += valEUR

		if x[0] == '3' && x[1] == '1' && x[4] == '1' {
			x[4] = '2'
			x[0] = '0'
			x[1] = '0'
		}
		if x[0] == '2' && x[1] == '8' && x[4] == '2' {
			x[4] = '3'
			x[0] = '0'
			x[1] = '0'
		}
		if x[0] == '3' && x[1] == '1' && x[4] == '3' {
			x[4] = '4'
			x[0] = '0'
			x[1] = '1'

		}

		if x[1] == '9' {
			x[0] += 1
			x[1] = '0'
		} else {
			x[1] += 1
		}
	}
	fmt.Println("")
	fmt.Println("USD max price:", USD, "date of max price:", MaxDateUSD, "avg USD price:", avgUSD/3)
	fmt.Println("EUR max price:", EUR, "date of max price:", MaxDateEUR, "avg EUR price:", avgEUR/3)
}
