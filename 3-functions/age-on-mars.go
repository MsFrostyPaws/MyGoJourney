/*
	How old would you be on Mars?
	A year on Earth has 365 days, while a year on Mars has 687 days.
*/

package main

import "fmt"

func mars_age(earth_age int) int {
	return earth_age * 365 / 687
}

func main() {
	var age int
	fmt.Scanln(&age)

	mars := mars_age(age)
	fmt.Println(mars)
}
