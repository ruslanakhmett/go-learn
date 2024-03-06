package main

import "fmt"

func main() {
	var x, p, y int

	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&y)

	years := 0
	for x < y {
		x = x + x*p/100
		years++
	}

	fmt.Println(years)
}