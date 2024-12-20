package main

func addStrings(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	res := make([]byte, len(num1)+1)
	i, j := len(num1)-1, len(num2)-1
	var inc byte
	for k := len(res) - 1; k >= 0; k-- {
		sum := inc
		if i >= 0 {
			sum += num1[i] - '0'
		}
		if j >= 0 {
			sum += num2[j] - '0'
		}
		res[k] = sum%10 + '0'
		inc = sum / 10
		i--
		j--
	}
	if res[0] > '0' {
		return string(res)
	} else {
		return string(res[1:])
	}
}

func main() {

}
