package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//"strconv"
	"slices"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		points := 0
		Nums := strings.Split(strings.Split(scanner.Text(), ":")[1], "|")
		winningNums := strings.Split(strings.Trim(Nums[0], " "), " ")
		myNums := strings.Split(strings.Trim(Nums[1], " "), " ")
		fmt.Print(winningNums)
		fmt.Print(" | ")
		fmt.Print(myNums)
		fmt.Print(" | ")

		for k := range myNums {
			if contains(winningNums, myNums[k]) {
				fmt.Print("FOUND: " + string(myNums[k]) + " | ")
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		fmt.Println(points)
		total += points
		//fmt.Println(winningNums)
		//fmt.Println(myNums)
	}
	fmt.Print("TOTAL: ")
	fmt.Println(total)
}

func contains(s []string, e string) bool {
    for _, a := range s {
		if a == "" || a == " " {
			continue
		}
        if a == e {
            return true
        }
    }
    return false
}