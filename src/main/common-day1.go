package main

func getLocal(c1 byte, c2 byte) int {
	if c1 != c2 {
		return 0
	}
	return (int)(c1 - 48)
}
