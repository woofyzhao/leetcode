package main

func isValid(s string) bool {
	var stack []byte
	for _, c := range []byte(s) {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			h := stack[len(stack)-1]
			if h == '(' && c == ')' || h == '{' && c == '}' || h == '[' && c == ']' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {

}
