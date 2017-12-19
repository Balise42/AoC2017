package main

import (
	"bufio"
	"day19"
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Can't open file", err)
	}

	maze := make([]string, 0, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		maze = append(maze, strings.TrimRight(scanner.Text(), "\n"))
	}

	var posX int
	var posY = 0
	for i := 0; i < len(maze[i]); i++ {
		if maze[0][i] == '|' {
			posX = i
			break
		}
	}

	direction := "down"
	step := 1

	for {
		if direction == "down" {
			posY = posY + 1
		} else if direction == "up" {
			posY = posY - 1
		} else if direction == "right" {
			posX = posX + 1
		} else if direction == "left" {
			posX = posX - 1
		}

		if posY < 0 || posY >= len(maze) || posX < 0 || posX > len(maze[posY]) {
			break
		}
		if maze[posY][posX] == ' ' {
			break
		}

		step = step + 1
		if maze[posY][posX] == '|' || maze[posY][posX] == '-' {
			continue
		} else if maze[posY][posX] == '+' {
			direction = day19.ChangeDirection(maze, direction, posY, posX)
		} else {
			fmt.Print(string(maze[posY][posX]))
		}
	}
	fmt.Println("")
	fmt.Println(step)
}
