package main

import "fmt"

func main() {
	var n, a int

	summ := 0
	fmt.Scan(&n)

	for i:=1; i<=n; i++ {
		fmt.Scan(&a)
		if (9 < a && a < 100) && (a % 8 == 0) {
			summ = summ + a
		}
	}
	fmt.Println(summ)
}