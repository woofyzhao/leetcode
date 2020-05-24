package main

const (
	S_Start = iota
	S_Up
	S_Down
)

func longestMountain(A []int) (result int) {
	if len(A) == 0 {
		return
	}
	A = append(A, A[len(A)-1])
	begin := 0
	state := S_Start
	for i := 1; i < len(A); i++ {
		switch state {
		case S_Start:
			if A[i] > A[i-1] {
				state = S_Up
				continue
			}
			begin = i
		case S_Up:
			if A[i] < A[i-1] {
				state = S_Down
			} else if A[i] == A[i-1] {
				state = S_Start
				begin = i
			}
		case S_Down:
			if A[i] >= A[i-1] {
				if i-begin > result {
					result = i - begin
				}
				if A[i] == A[i-1] {
					state = S_Start
					begin = i
				} else {
					state = S_Up
					begin = i - 1
				}
			}
		}
	}
	return
}

func main() {
	return
}

// #数组 #状态机
// #array #state_machine
