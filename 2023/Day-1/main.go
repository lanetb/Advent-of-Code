package main

import (
    "fmt"
	//"io"
    "os"
	"bufio"
    "unicode"
    "strconv"
)

func main() {
	file, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }

	defer file.Close()
    total := 0
	scanner := bufio.NewScanner(file)
    
    for scanner.Scan() {
        digit := ""
        for _, c := range scanner.Text() {
            if unicode.IsDigit(c){
                digit = digit + string(c)
                break
            }
        }
        for i := len(scanner.Text()) - 1; i >= 0; i-- {
            c := scanner.Text()[i]
            if unicode.IsDigit(rune(c)) {
                digit = digit + string(c)
                break
            }
        }
        fmt.Println(digit)
        i, _ := strconv.Atoi(digit)
        total += int(i)
    }
    fmt.Printf("TOTAL = %d", total)
 
    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
    // make a buffer to keep chunks that are read
	
}