package country

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestSanity(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(123, 123, "they should be equal")
}

func TestCsvCountry(t *testing.T) {
	csvStr := `
Country,Date,Age,Amount,Id
"United States",2012-02-01,50,112.1,01234
"United States",2012-02-01,32,321.31,54320
"United Kingdom",2012-02-01,17,18.2,12345
"United States",2012-02-01,32,321.31,54320
"United Kingdom",2012-02-01,NA,18.2,12345
"United States",2012-02-01,32,321.31,54320
"United States",2012-02-01,32,321.31,54320
Spain,2012-02-01,66,555.42,00241
`
	df := dataframe.ReadCSV(strings.NewReader(csvStr))
	fmt.Println(df)
}
func TestCountryJson(t *testing.T) {
	df := LoadCountryDataFrame()
	fmt.Println("[[[TestCountryJson]]]", df)
}

func TestTsv(t *testing.T) {
	// Using with a string reader
	csvString := "ISO\tISO3\tISO-Numeric\tfips\tCountry\tCapital\tArea(in sq km)\tPopulation\tContinent\ttld\tCurrencyCode\tCurrencyName\tPhone\tPostal Code Format\tPostal Code Regex\tLanguages\tgeonameid\tneighbours\tEquivalentFipsCode\nAD\tAND\t020\tAN\tAndorra\tAndorra la Vella\t468\t77006\tEU\t.ad\tEUR\tEuro\t376\tAD###\t^(?:AD)*(\\d{3})$\tca\t3041565\tES,FR"
	df := dataframe.ReadCSV(strings.NewReader(csvString))
	fmt.Println("TestTsv", df)
}

//func TestAdd(t *testing.T) {
//	assert.Equal(t, 123, 123, "they should be equal")
//	//testCases := []struct {
//	//	a        int
//	//	b        int
//	//	expected int
//	//}{
//	//	{2, 3, 5},
//	//	{-1, -1, -2},
//	//	{0, 0, 0},
//	//	{100, 200, 300},
//	//}
//	//
//	//for _, tc := range testCases {
//	//	result := Add(tc.a, tc.b)
//	//	if result != tc.expected {
//	//		t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
//	//	}
//	//}
//}
