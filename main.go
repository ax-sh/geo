package main

import (
	_ "embed"
	"geo/cmd"
)

//go:embed VERSION
var version string

func main() {
	cmd.Execute(version)

	//s := "gopher"
	//fmt.Printf("Hello and welcome, %s! version [%s]\n", s, version)
	//countryCode := "41"
	//

	////sel := fil.Select([]string{"ISO", "ISO3", "ISO-Numeric", "fips",
	////	"Country",
	////	"Capital",
	////	//,
	////	"Population",
	////	//"Continent	tld	CurrencyCode",
	////	"CurrencyName",
	////	"Phone",
	////	//"Postal Code Format	Postal Code Regex",
	////	"Languages",
	////	"geonameid",
	////	"neighbours",
	////	//"EquivalentFipsCode",
	////})

	//fmt.Println(result)

}
