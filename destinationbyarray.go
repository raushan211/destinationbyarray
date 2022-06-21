package main

import "fmt"

var arr = [][]string{{"Lima", "London"}, {"New York", "Sao Paulo"}, {"London", "New York"}}

func main() {
	allCities := []string{}
	for i := 0; i < len(arr); i++ {
		allCities = append(allCities, arr[i][0])
		allCities = append(allCities, arr[i][1])
	}
	//allCities = {"Lima", "London", "New York", "Sao Paulo", "London", "New York"}

	for i := 0; i < len(arr); i++ {
		//cancel source cities which are at arr[i][0]
		allCities = deleteCity(allCities, arr[i][0])
	}
	fmt.Println(allCities)

}

func deleteCity(arr []string, element string) []string {
	res := []string{}
	for _, v := range arr {
		if v != element {
			res = append(res, v)
		}
	}
	return res
}
