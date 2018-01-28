package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for num := range nums {
		sum += num
	}
	fmt.Println(sum)
}
