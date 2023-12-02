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

	GAMEREGEX, _ := regexp.Compile("Game [0-9]+")
	AMOUNTREGEX, _ := regexp.Compile("[1-9]+ [a-z]+")
	currentBag := Bag{red: 12, green: 13, blue: 14}

	file, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }

	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
    for scanner.Scan() {
		valid_game := true
		gamenumber := strings.Split(string(GAMEREGEX.Find(scanner.Bytes())), " ")[1]
		amounts := AMOUNTREGEX.FindAll(scanner.Bytes(), -1)

		for _, item := range amounts {
			items := strings.Split(string(item), " ")
			switch items[1] {
			case "red":
				itemstr := string(items[0])
				i, _:= strconv.Atoi(itemstr)
				if i > currentBag.red {
					valid_game = false
				}
			case "green":
				itemstr := string(items[0])
				i, _:= strconv.Atoi(itemstr)
				if i > currentBag.green {
					valid_game = false
				}
			case "blue":
				itemstr := string(items[0])
				i, _:= strconv.Atoi(itemstr)
				if i > currentBag.blue {
					valid_game = false
				}
			}
		}

		if valid_game {
			num, _ := strconv.Atoi(gamenumber)
			total += num
		}

	}

	fmt.Printf("Total: %d\n", total)
}