package main

import "fmt"

func partitionLabels(S string) (result []int) {
	n := len(S)
	endPos := make(map[uint8]int, n)
	for i := 0; i < n; i++ {
		endPos[S[i]] = i
	}
	next := 0
	len := 0
	for i := 0; i < n; i++ {
		len += 1
		if endPos[S[i]] > next {
			next = endPos[S[i]]
		}
		if i == next {
			result = append(result, len)
			len = 0
		}
	}
	return
}

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println(partitionLabels("1"))
	fmt.Println(partitionLabels("ababcbaxcadefegdehijxhwwvvsy"))
}
