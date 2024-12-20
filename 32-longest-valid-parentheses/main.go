package main

import "fmt"

// stack
func longestValidParentheses0(s string) int {
	stack, res := []int{}, 0
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				res = max(res, i-stack[len(stack)-1])
			}
		}
	}
	return res
}

// dp: dp[i] = dp[i-2] + 2 || dp[i - dp[i-1] - 2] + dp[i-1] + 2
func longestValidParentheses(s string) (res int) {
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == '(' {
			continue
		}
		if s[i-1] == '(' {
			dp[i] = 2
			if i-2 >= 0 {
				dp[i] += dp[i-2]
			}
		} else if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
			dp[i] = dp[i-1] + 2
			if i-dp[i-1]-2 >= 0 {
				dp[i] += dp[i-dp[i-1]-2]
			}
		}
		res = max(res, dp[i])
	}
	return
}

func longestValidParentheses2(s string) (res int) {
	stack := make([]int, 0)
	start := -1
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
			continue
		}
		if len(stack) == 0 {
			start = i
			continue
		}
		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			res = max(res, i-start)
		} else {
			res = max(res, i-stack[len(stack)-1])
		}
	}
	return
}

func main() {
	fmt.Println(longestValidParentheses2(")()())"))
	fmt.Println(longestValidParentheses2("(()()"))
}
