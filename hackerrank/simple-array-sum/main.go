package main

import "fmt"

func simpleArraySum(nums []int) {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println(sum)
}

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	simpleArraySum(a)
}
