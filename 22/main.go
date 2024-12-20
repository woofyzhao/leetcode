package main

func generateParenthesis(n int) (res []string) {
	var gen func(int, int, string)
	gen = func(n int, stack int, parens string) {
		if n == 0 && stack == 0 {
			res = append(res, parens)
			return
		}
		if n > 0 {
			gen(n-1, stack+1, parens+"(")
		}
		if stack > 0 {
			gen(n, stack-1, parens+")")
		}
	}
	gen(n, 0, "")
	return
}

func main() {

}
