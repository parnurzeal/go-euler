package main

import "fmt"

func main() {
	const input = 200
	coinTypes := []int{1, 2, 5, 10, 20, 50, 100, 200}
	ways := make([]int, 201)
	ways[0] = 1
	fmt.Println(ways)
	for _, i := range coinTypes {
		for j := i; j <= input; j++ {
			ways[j] += ways[j-i]
		}
		fmt.Println(ways)
	}
	fmt.Println(ways[200])
}
