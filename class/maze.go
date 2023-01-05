package class

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type MazeClass struct {
	contents    [][]string
	rows        []string
	height      int
	width       int
	walls       [][]bool
	start       [2]int
	goal        [2]int
	solution    [][2]int
	numExplored int
	explored    [][]int
}

func ArrayContent(content string) [][]string {
	rows := strings.Split(content, "\n")
	var res [][]string
	for _, r := range rows {
		res = append(res, strings.Split(r, ""))
	}
	return res
}

func InitMaze(filename string) MazeClass {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	text := string(file)
	rows := strings.Split(text, "\n")
	height := len(rows)
	width := len(rows[0])
	contents := ArrayContent(text)
	var solution [][2]int
	var start [2]int
	var walls [][]bool
	var goal [2]int

	for i, _ := range rows {
		var row []bool
		for j, _ := range rows[i] {
			if contents[i][j] == "A" {
				start = [2]int{i, j}
				row = append(row, false)
			} else if contents[i][j] == "B" {
				goal = [2]int{i, j}
				row = append(row, false)
			} else if contents[i][j] == " " {
				row = append(row, false)
			} else {
				row = append(row, true)
			}
		}
		walls = append(walls, row)
	}

	return MazeClass{contents: contents, rows: rows, height: height, width: width, walls: walls, start: start, goal: goal, solution: solution}
}

type EnumratedBool struct {
	index int
	value []bool
}
type EnumratedEnum struct {
	index int
	value EnumratedBool
}

func EnumrateBool(value [][]bool) []EnumratedBool {
	var res []EnumratedBool
	for i, w := range value {
		arr := EnumratedBool{i, w}
		res = append(res, arr)
	}

	return res
}

func EnumrateEnum(value []EnumratedBool) []EnumratedEnum {
	var res []EnumratedEnum
	for i, w := range value {
		arr := EnumratedEnum{i, w}
		res = append(res, arr)
	}

	return res
}
func isElementExist(s [][2]int, value [2]int) bool {
	for _, v := range s {
		if v[0] == value[0] && v[1] == value[1] {
			return true
		}
	}
	return false
}

func (maze MazeClass) Print() {
	var solution [][2]int
	if len(maze.solution) > 0 {
		solution = maze.solution
	}
	for i, row := 0, EnumrateBool(maze.walls); i < len(row); i++ {
		for j, col := 0, EnumrateEnum(row); j <= len(col); j++ {
			value := [2]int{i, j}
			if col[i].value.value[j] {
				fmt.Print("█")
			} else if maze.start[0] == i && maze.start[1] == j {
				fmt.Print("A")
			} else if maze.goal[0] == i && maze.goal[1] == j {
				fmt.Print("B")
			} else if isElementExist(solution, value) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func (maze MazeClass) Neighbours(state any) {
}

var (
	frontier      StackFrontierClass
	queueFrontier QueueFrontierClass
)

func ReverseInts(input [2]int) []int {
	var output []int

	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}

	return output
}

func ReverseInt2(input [][2]int) [][2]int {
	var output [][2]int

	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}

	return output
}

func (maze MazeClass) Solve() {
	maze.numExplored = 0

	var parent *NodeClass
	var action [2]int

	start := NodeClass{state: maze.start, parent: parent, action: action}
	frontier = StackFrontierClass{}
	frontier.Add(start)

	for 0 == 0 {
		if frontier.Empty() {
			panic(errors.New("no solution"))
		}

		node := frontier.Remove()
		maze.numExplored++

		if node.state == maze.goal {
			var actions [2]int
			var cells [2]int

			for len(node.parent.state) > 0 {
				actions = append(actions, node.action)
				cells = append(cells, node.state)
				node = *node.parent
			}
			actions = ReverseInts(actions)
			cells = ReverseInts(cells)
			maze.solution = [][2]int{actions, cells}
		}
	}
}
