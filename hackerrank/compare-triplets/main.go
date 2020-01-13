package main

import "fmt"

func compareTriplet(alArr []int, boArr []int) {
	alice := 0
	bob := 0

	for i := 0; i < 3; i++ {
		if alArr[i] > boArr[i] {
			alice++
		} else if boArr[i] > alArr[i] {
			bob++
		}
	}

	fmt.Println(alice, bob)
}

func main() {
	a := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Scan(&a[i])
	}

	b := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Scan(&b[i])
	}

	compareTriplet(a, b)
}
