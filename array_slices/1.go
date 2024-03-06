package main

import "fmt"

func main() {
	var (
	workArray [10]uint16
	i1, i2, i3, i4, i5, i6 uint16
	)

	for idx := range workArray {
		fmt.Scan(&workArray[idx])
	}
	fmt.Scan(&i1, &i2, &i3, &i4, &i5, &i6)

	workArray[i1], workArray[i2] = workArray[i2], workArray[i1]
	workArray[i3], workArray[i4] = workArray[i4], workArray[i3]
	workArray[i5], workArray[i6] = workArray[i6], workArray[i5]

	for idx := range workArray {
		fmt.Print(workArray[idx], " ")
	}
}