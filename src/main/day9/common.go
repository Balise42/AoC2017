package day9

func CleanData(data []byte) ([]byte, int) {
	inGarbage := false
	ignored := false
	garbage := 0
	res := make([]byte, 0, len(data))
	for _, c := range data {
		if ignored {
			ignored = false
			continue
		} else {
			if inGarbage {
				if c == '>' {
					inGarbage = false
					continue
				} else if c == '!' {
					ignored = true
					continue
				}
				garbage = garbage + 1
			} else {
				if c == '<' {
					inGarbage = true
					continue
				}
			}
			if !inGarbage {
				res = append(res, c)
			}
		}
	}
	return res, garbage
}

func ComputeScore(data []byte) int {
	level := 0
	res := 0
	for _, c := range data {
		if c == '{' {
			level = level + 1
		} else if c == '}' {
			res = res + level
			level = level - 1
		}
	}
	return res
}
