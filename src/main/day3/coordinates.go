package day3

import (
	"math"
)

func GetCoordinates(location int) (float64, float64) {
	loc64 := float64(location)
	lowsqrt := math.Floor(math.Sqrt(loc64))
	highsqrt := math.Ceil(math.Sqrt(loc64))

	if lowsqrt*lowsqrt == loc64 && location%2 == 1 {
		return (lowsqrt - 1) / 2, -(lowsqrt - 1) / 2
	}

	if lowsqrt*lowsqrt == loc64 && location%2 == 0 {
		return -(lowsqrt/2 - 1), lowsqrt / 2
	}

	if lowsqrt * lowsqrt + 1 == loc64 && location % 2 == 1 {
		return -lowsqrt/2, lowsqrt / 2
	}


	if lowsqrt * lowsqrt + 1 == loc64 && location % 2 == 0 {
		return (lowsqrt - 1) / 2 + 1, -(lowsqrt - 1) / 2
	}

	mid := math.Ceil(((lowsqrt * lowsqrt) + (highsqrt * highsqrt)) / 2)

	if int(highsqrt) % 2 == 1 {
		tl := lowsqrt * lowsqrt + 1
		br := highsqrt * highsqrt

		if loc64 <= mid {
			x, y := GetCoordinates(int(tl))
			return x, y-(loc64 - tl)
		} else {
			x, y := GetCoordinates(int(br))
			return x-(br - loc64), y
		}
	}


	br := lowsqrt * lowsqrt + 1
	tl := highsqrt * highsqrt +1

	if loc64 > mid {
		x, y := GetCoordinates(int(tl))
		return x+(tl-loc64), y
	}
	x, y := GetCoordinates(int(br))
	return x, y+(loc64-br)

}

