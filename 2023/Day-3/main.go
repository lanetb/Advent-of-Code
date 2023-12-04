package main

import (
	"fmt"
	"os"
	"bufio"
)

type number struct {
	val 	string
	xStart 	int
	xEnd 	int
	y 		int
}
type Pair struct {
	x int
	y int
}


func main() {
	NUMBERS := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	file, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
	maap := make([][]string, 141)
	scanner := bufio.NewScanner(file)
	currline := 0
	for scanner.Scan() {
		currline += 1
		maap[currline] = make([]string, 140)
		for i, c := range scanner.Text() {
			maap[currline][i] = string(c)
		}
	}
}

func findNumbers(maap [][]string) []number {
	ROW := len(maap)
	COL := len(maap[0])
	visit := make([][]bool, ROW)
	for i := range visit {
		visit[i] = make([]bool, COL)
		for j := range visit[i] {
			visit[i][j] = false
		}
	}
	numbers := []number{}

	directions := []Pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {1, -1}, {-1, 1}, {1, 1}, {-1, -1}}
	var bfs = func (row int, col int) {
		queue := []Pair{{row, col}}
		visit[row][col] = true
		for len(queue) != 0 {
			curr := queue[0]
			queue = queue[1:]
			for _, dir := range directions {
				newRow := curr.x + dir.x
				newCol := curr.y + dir.y
				if newRow >= 0 && newRow < ROW && newCol >= 0 && newCol < COL && maap[newRow][newCol] != " " && !visit[newRow][newCol] {
					visit[newRow][newCol] = true
					queue = append(queue, Pair{newRow, newCol})
				}
			}
		}
	}
}
func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}




