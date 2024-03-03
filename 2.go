package main

import "fmt"

func main() {
	var (
		a int
		counter int
		max int
	)
	fmt.Scan(&a)
	max = a
	counter = 1

	for a != 0 {
		fmt.Scan(&a)
		if a > max {
			max = a
			counter = 1
		} else if a == max {
			counter++
		}
	}
	fmt.Println(counter)
}