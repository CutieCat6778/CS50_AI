package main

import "cutiecat6778/cs05_ai/class"

var (
	Maze          class.MazeClass
	StackFrontier class.StackFrontierClass
	QueueFrontier class.QueueFrontierClass
)

func init() {
	Maze = class.InitMaze("maze.txt")
}

func main() {
	Maze.Print()
}
