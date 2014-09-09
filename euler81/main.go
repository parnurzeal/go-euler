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
		nums := strings.Split(oneLine, " ")
		matrix[r] = make([]int, len(nums))
		for c, oneNum := range nums {
			matrix[r][c], err = strconv.Atoi(oneNum)
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Println(matrix)
}

type Graph struct {
	nodes Node
}
type Node struct {
	payload int
	adjNode map[int]struct{}
}

type Queue []int

func (q *Queue) enqueue(elem int) {
	*q = append(*q, elem)
}

func (q *Queue) dequeue() int {
	if q.isEmpty() {
		panic("try to dequeue empty queue")
	}
	elem := (*q)[0]
	copy((*q)[0:], (*q)[1:])
	(*q)[len(*q)-1] = 0
	(*q) = (*q)[:len(*q)-1]
	return elem
}

func (q *Queue) isEmpty() bool {
	if len(*q) == 0 {
		return true
	} else {
		return false
	}
}
