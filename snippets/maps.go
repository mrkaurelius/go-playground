package main

import (
	"fmt"
)

type Country struct {
	Abbrevation string
	Fullname    string
}

func main() {
	fmt.Println("Merhaba yalan dunya")

	countryMap := make(map[string]string)
	countryMap["USA"] = "United states of America"
	countryMap["TR"] = "Turkey"

	for key := range countryMap {
		fmt.Println(key)
	}

	countryMapStruct := make(map[string]Country)
	countryMapStruct["tr"] = Country{Abbrevation: "TR", Fullname: "Turkey"}
	countryMapStruct["usa"] = Country{Abbrevation: "USA", Fullname: "United States of America"}

	for key, value := range countryMapStruct {
		fmt.Println(key, value)
	}

	str := "merhaba yalan dunya"

	for i, v := range str {
		fmt.Println(i, string(v))
	}

}
