package day19

func ChangeDirection(maze []string, direction string, posY int, posX int) string {
	if direction == "left" || direction == "right" {
		if posY >= len(maze)-1 || maze[posY+1][posX] == ' ' {
			return "up"
		}
		return "down"
	} else {
		if posX >= len(maze[0])-1 || maze[posY][posX+1] == ' ' {
			return "left"
		}
		return "right"
	}
}
