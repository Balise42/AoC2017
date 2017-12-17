package main

import (
	"./day3"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	location, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Input must be an integer.")
	}

	x, y := day3.GetCoordinates(location)
	fmt.Println(x, y)
	fmt.Println(math.Abs(x) + math.Abs(y))
}
