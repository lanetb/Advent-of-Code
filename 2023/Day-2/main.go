package main

import (
	"fmt"
	//"io"
	"os"
	"bufio"
	//"unicode"
	"strconv"
	"regexp"
	"strings"
)

func main() {
	type Bag struct {
		red   int
		green int
		blue  int
	}

	//GAMEREGEX, _ := regexp.Compile("Game [0-9]+")
	AMOUNTREGEX, _ := regexp.Compile("[0-9]+ [a-z]+")
	//currentBag := Bag{red: 12, green: 13, blue: 14}

	file, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }

	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
    for scanner.Scan() {
		//valid_game := true
		//gamenumber := strings.Split(string(GAMEREGEX.Find(scanner.Bytes())), " ")[1]
		amounts := AMOUNTREGEX.FindAll(scanner.Bytes(), -1)
		minred, mingreen, minblue := 0, 0, 0

		for _, item := range amounts {
			items := strings.Split(string(item), " ")
			switch items[1] {
			case "red":
				itemstr := string(items[0])
				i, _:= strconv.Atoi(itemstr)
				fmt.Printf("Red: %d\n", i)
				if i > minred {
					minred = i
				}
			case "green":
				itemstr := string(items[0])
				fmt.Printf("Green: %s\n", itemstr)
				i, _:= strconv.Atoi(itemstr)
				if i > mingreen {
					mingreen = i
				}
			case "blue":
				itemstr := string(items[0])
				fmt.Printf("Blue: %s\n", itemstr)
				i, _:= strconv.Atoi(itemstr)
				if i > minblue {
					minblue = i
				}
			}
		}
		fmt.Printf("Red: %d, Green: %d, Blue: %d\n", minred, mingreen, minblue)
		temptotal := minred * mingreen * minblue
		fmt.Printf("TempTotal: %d\n", temptotal)
		total += temptotal
	}

	fmt.Printf("Total: %d\n", total)
}