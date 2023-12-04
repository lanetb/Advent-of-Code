package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//"strconv"
	//"slices"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	total := 0
	text := []string{} // Declare the text variable as an empty string slice

	for scanner.Scan() {
		text = append(text, scanner.Text()) // Append the scanned text to the text slice
	}
	for t, _ := range text {
		// Your code here
		points := 0
		Nums := strings.Split(strings.Split(text[t], ":")[1], "|")
		winningNums := strings.Split(strings.Trim(Nums[0], " "), " ")
		myNums := strings.Split(strings.Trim(Nums[1], " "), " ")
		fmt.Print(winningNums)
		fmt.Print(" | ")
		fmt.Print(myNums)
		fmt.Print(" | ")

		for k := range myNums {
			if contains(winningNums, myNums[k]) {
				fmt.Print("FOUND: " + string(myNums[k]) + " | ")
				points = 1
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