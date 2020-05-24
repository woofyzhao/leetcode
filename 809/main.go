package main

func expressiveWords(S string, words []string) (result int) {
	for _, from := range words {
		f := 0
		t := 0
		strechy := true
		for f < len(from) && t < len(S) {
			if from[f] != S[t] {
				strechy = false
				break
			}
			fn := 1
			tn := 1
			for f < len(from)-1 && from[f+1] == from[f] {
				f += 1
				fn += 1
			}
			for t < len(S)-1 && S[t+1] == S[t] {
				t += 1
				tn += 1
			}
			if tn < 3 && fn != tn {
				strechy = false
				break
			}
			if tn >= 3 && fn > tn {
				strechy = false
				break
			}
			f += 1
			t += 1
		}
		if strechy && (f == len(from)) && (t == len(S)) {
			result += 1
		}
	}
	return
}

func main() {

}
