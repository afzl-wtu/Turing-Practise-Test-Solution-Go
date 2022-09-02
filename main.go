package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter your score sperated by spaces and in capital letters:")
	reader := bufio.NewReader(os.Stdin)
	rawInput, _ := reader.ReadString('\r')
	rawInput = strings.Replace(rawInput, "\r", " ", -1)
	ops := strings.Split(rawInput, " ")
	fmt.Println(calPoints(ops))
}

func calPoints(ops []string) int {
	var res int = 0
	var totalSlice = []int{}
	for _, value := range ops {
		val, er := strconv.ParseInt(value, 10, 32)
		if er != nil {
			var tempTotalSlice []int
			switch value {
			case "C":
				for i := 0; i < (len(totalSlice) - 1); i++ {
					tempTotalSlice = append(tempTotalSlice, totalSlice[i])

				}
				totalSlice = tempTotalSlice
			case "D":
				totalSlice = append(totalSlice, 2*totalSlice[len(totalSlice)-1])
			case "+":
				totalSlice = append(totalSlice, totalSlice[len(totalSlice)-1]+totalSlice[len(totalSlice)-2])
			default:

			}
		}
		if er == nil {
			totalSlice = append(totalSlice, int(val))
		}
	}
	for _, val := range totalSlice {
		res += val
	}
	return res
}
