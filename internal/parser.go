package internal

import (
	model "github.com/LittleMikle/html_parser_go/pkg"
	"github.com/nfx/go-htmltable"
	"strconv"
	"strings"
)

var USD, avgUSD float64 = 0, 0
var EUR, avgEUR float64 = 0, 0

func DataParser(currentDate string) (float64, float64) {
	url := "https://www.cbr.ru/currency_base/daily/?UniDbQuery.Posted=True&UniDbQuery.To=" + currentDate

	answer, err := htmltable.NewSliceFromURL[model.Money](url)
	if err != nil {

	}

	valuteUSD := strings.Replace(answer[13].Value, ",", ".", -1)
	valuteEUR := strings.Replace(answer[14].Value, ",", ".", -1)

	valUSD, err := strconv.ParseFloat(valuteUSD, 64)
	valEUR, err := strconv.ParseFloat(valuteEUR, 64)

	return valUSD, valEUR

}
