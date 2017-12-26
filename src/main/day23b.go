package main

import "fmt"

func main() {
	a := 1
	b := 81
	c := 81
	h := 0
	if a != 0 {
		goto L1
	}
	goto start
L1:
	b = b * 100
	b = b + 100000
	c = b
	c = c + 17000
start:
	f := 1
	d := 2
L5:
	e := 2

L4:
	g := d
	g = g * e
	g = g - b
	if g != 0 {
		goto L3
	}
	f = 0
L3:
	e = e + 1
	g = e
	g = g - b
	if g != 0 {
		goto L4
	}
	d = d + 1
	g = d
	g = g - b
	if g != 0 {
		goto L5
	}
	if f != 0 {
		goto L6
	}
	h = h + 1
L6:
	g = b
	g = g - c
	if g != 0 {
		goto L7
	}
	goto end
L7:
	b = b + 17
	goto start
end:
	fmt.Println(h)
}