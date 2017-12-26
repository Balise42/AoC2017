package main

import "fmt"

func main() {

	b := 81
	c := 81
	h := 0

	b = b * 100
	b = b + 100000
	c = b
	c = c + 17000
	i := 0
	for {
		i++
		f := 1
		d := 2

		g := 1
		e:=2
		for g != 0 {
			e = 2

			for g != 0 {
				g = d
				g = g * e
				g = g - b
				if g == 0 {
					f = 0
					goto L
				}

				e = e + 1
				g = e
				g = g - b
			}
			d = d + 1
			g = d
			g = g - b
		}
		fmt.Println(d)
		L: if f == 0 {
			h = h + 1
			fmt.Println(i, "inc")
		} else {
			fmt.Println(i, "nop")
		}

		g = b
		g = g - c
		if g == 0 {
			break
		}
		fmt.Println(h, g, b, c)

		b = b + 17
	}

	fmt.Println(h)
}
