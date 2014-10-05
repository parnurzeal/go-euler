package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	inBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(inBytes), "\n")
	// put in matrix
	matrix := make([][]int, len(lines))
	for r, oneLine := range lines {
		nums := strings.Split(oneLine, ",")
		fmt.Println(nums)
		matrix[r] = make([]int, len(nums))
		for c, oneNum := range nums {
			matrix[r][c], err = strconv.Atoi(oneNum)
			if err != nil {
				panic(err)
			}
		}
	}
	//fmt.Println(matrix)
	dyna := make([][]int, len(matrix))
	for r, _ := range matrix {
		dyna[r] = make([]int, len(matrix[r]))
	}
	dyna[0][0] = matrix[0][0]
	for c := 1; c < len(matrix[0]); c++ {
		dyna[0][c] += dyna[0][c-1] + matrix[0][c]
	}
	for r := 1; r < len(matrix); r++ {
		dyna[r][0] += dyna[r-1][0] + matrix[r][0]
	}
	for r := 1; r < len(matrix); r++ {
		for c := 1; c < len(matrix[r]); c++ {
			sum1 := matrix[r][c] + dyna[r-1][c]
			sum2 := matrix[r][c] + dyna[r][c-1]
			if sum1 < sum2 {
				dyna[r][c] = sum1
			} else {
				dyna[r][c] = sum2
			}
		}
	}
	fmt.Println(dyna[len(dyna)-1][len(dyna[0])-1])
}
