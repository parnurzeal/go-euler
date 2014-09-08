package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	inBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(inBytes))
	lines := strings.Split(string(inBytes), "\n")
	// put in matrix
	for idx, oneLine := range lines {
		fmt.Println(idx, oneLine)
		/*nums := strings.Split(oneLine, " ")
		if len(nums) == 0 {
			break
		}
		fmt.Println(len(nums), nums)*/
	}

	fmt.Println(lines)
}
