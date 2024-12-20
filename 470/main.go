package main

func rand7() int {
	return 1
}

func rand10() int {
	n := 49
	for n >= 40 { //0-39
		a := rand7()
		b := rand7()
		n = (a-1)*7 + b - 1 //0-48
	}
	return n%10 + 1
}

func main() {

}
