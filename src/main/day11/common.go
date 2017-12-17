package day11

func CancelCompletely(tally map[string]int, k1 string, k2 string) {
	if tally[k1] > tally[k2] {
		tally[k1] = tally[k1] - tally[k2]
		tally[k2] = 0
	} else {
		tally[k2] = tally[k2] - tally[k1]
		tally[k1] = 0
	}
}

func CancelToNS(tally map[string]int, k1 string, k2 string, k3 string) {
	if tally[k1] > tally[k2] {
		tally[k1] = tally[k1] - tally[k2]
		tally[k3] = tally[k2]
		tally[k2] = 0
	} else {
		tally[k2] = tally[k2] - tally[k1]
		tally[k3] = tally[k1]
		tally[k1] = 0
	}
}
