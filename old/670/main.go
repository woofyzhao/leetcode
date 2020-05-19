package main

import (
	"fmt"
	"strconv"
)

func maximumSwap(num int) int {
	bytes := []byte(strconv.Itoa(num))
	maxd := byte(0)
	maxi := -1
	swapfrom := -1
	swapto := -1
	for i := len(bytes) - 1; i >= 0; i-- {
		if bytes[i] > maxd {
			maxd = bytes[i]
			maxi = i
		} else if bytes[i] < maxd {
			//fmt.Println("-", i)
			swapfrom = maxi
			swapto = i
		}
	}
	if swapfrom >= 0 && swapto >= 0 {
		bytes[swapfrom], bytes[swapto] = bytes[swapto], bytes[swapfrom]
	}
	res, _ := strconv.Atoi(string(bytes))
	return res
}

func main() {
	fmt.Println(maximumSwap(2736))
	fmt.Println(maximumSwap(9973))
	fmt.Println(maximumSwap(9937))
	fmt.Println(maximumSwap(100001))
	fmt.Println(maximumSwap(0))
	fmt.Println(maximumSwap(12))
	fmt.Println(maximumSwap(333))
	fmt.Println(maximumSwap(12909))
	fmt.Println(maximumSwap(1290912))
}
