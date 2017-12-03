package main

import (
	"os"
	"math"
	"fmt"
	"strconv"
	"log"
	"./day3"
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