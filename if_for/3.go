package main

import "fmt"

func main() {
	var 
		n, c, d, i int

	fmt.Scan(&n)
	fmt.Scan(&c)
	fmt.Scan(&d)

	for i=0; i<=n; i++ {
		if i % c == 0 && i % d != 0 {
			fmt.Println(i)
			break
		}
	}
}