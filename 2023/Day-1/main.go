package main

import (
    "fmt"
	//"io"
    "os"
	"bufio"
    "unicode"
    "strconv"
)

var ONE string = "one"
var TWO string = "two"
var THREE string = "three"
var FOUR string = "four"
var FIVE string = "five"
var SIX string = "six"
var SEVEN string = "seven"
var EIGHT string = "eight"
var NINE string = "nine"

func main() {
	file, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }

	defer file.Close()
    total := 0
	scanner := bufio.NewScanner(file)
    
    for scanner.Scan() {
        fmt.Printf("%s ", scanner.Text())
        digit := ""
        forloop:for i, c := range scanner.Text() {
            if unicode.IsDigit(c){
                digit = digit + string(c)
                break
            } else{
                switch byte(c) {
                    case ONE[0]:
                        if len(scanner.Text()[i:]) >= 3 {
                            if scanner.Text()[i:i+3] == ONE {
                                digit = digit + "1"
                                break forloop
                            }
                        } 
                    case TWO[0]: 
                        if len(scanner.Text()[i:]) >= 3 {
                            if scanner.Text()[i:i+3] == TWO {
                                digit = digit + "2"
                                break forloop
                            }
                        }
                        fallthrough
                    case THREE[0]:
                        if len(scanner.Text()[i:]) >= 5 {
                            if scanner.Text()[i:i+5] == THREE {
                                digit = digit + "3"
                                break forloop
                            }
                        } 
                    case FOUR[0]: 
                        if len(scanner.Text()[i:]) >= 4 {
                            if scanner.Text()[i:i+4] == FOUR {
                                digit = digit + "4"
                                break forloop
                            } 
                        }
                        fallthrough
                    case FIVE[0]:
                        if len(scanner.Text()[i:]) >= 4 {
                            if scanner.Text()[i:i+4] == FIVE {
                                digit = digit + "5"
                                break forloop
                            }
                        } 
                    case SIX[0]: 
                        if len(scanner.Text()[i:]) >= 3 {
                            if scanner.Text()[i:i+3] == SIX {
                                digit = digit + "6"
                                break forloop
                            } 
                        } 
                        fallthrough
                    case SEVEN[0]:
                        if len(scanner.Text()[i:]) >= 5 {
                            if scanner.Text()[i:i+5] == SEVEN {
                                digit = digit + "7"
                                break forloop
                            }
                        }
                    case EIGHT[0]:
                        if len(scanner.Text()[i:]) >= 5 {
                            if scanner.Text()[i:i+5] == EIGHT {
                                digit = digit + "8"
                                break forloop
                            }
                        }
                    case NINE[0]:
                        if len(scanner.Text()[i:]) >= 4 {
                            if scanner.Text()[i:i+4] == NINE {
                                digit = digit + "9"
                                break forloop
                            }
                        }
                }
            }
        }

        forloop2:for i := len(scanner.Text()) - 1; i >= 0; i-- {
            c := scanner.Text()[i]
            if unicode.IsDigit(rune(c)) {
                digit = digit + string(c)
                break
            } else {
                switch byte(c) {
                    case ONE[2]: 
                        if len(scanner.Text()[:i]) >= 3 {
                            if scanner.Text()[i-2:i+1] == ONE {
                                digit = digit + "1"
                                break forloop2
                            }
                        } 
                        fallthrough
                    case THREE[4]:
                        if len(scanner.Text()[:i]) >= 5 {
                            if scanner.Text()[i-4:i+1] == THREE {
                                digit = digit + "3"
                                break forloop2
                            }
                        }
                        fallthrough
                    case FIVE[3]:
                        if len(scanner.Text()[:i]) >= 4 {
                            if scanner.Text()[i-3:i+1] == FIVE {
                                digit = digit + "5"
                                break forloop2
                            }
                        }
                        fallthrough
                    case NINE[3]:
                        if len(scanner.Text()[:i]) >= 4 {
                            if scanner.Text()[i-3:i+1] == NINE {
                                digit = digit + "9"
                                break forloop2
                            }
                        }

                    case TWO[2]:
                        if len(scanner.Text()[:i]) >= 2 {
                            if scanner.Text()[i-2:i+1] == TWO {
                                digit = digit + "2"
                                break forloop2
                            }
                        }  
                    case FOUR[3]: 
                        if len(scanner.Text()[:i]) >= 4 {
                            if scanner.Text()[i-3:i+1] == FOUR {
                                digit = digit + "4"
                                break forloop2
                            } 
                        } 
                    case SIX[2]: 
                        if len(scanner.Text()[:i]) >= 3 {
                            if scanner.Text()[i-2:i+1] == SIX {
                                digit = digit + "6"
                                break forloop2
                            } 
                        } 
                    case SEVEN[4]:
                        if len(scanner.Text()[:i]) >= 5 {
                            if scanner.Text()[i-4:i+1] == SEVEN {
                                digit = digit + "7"
                                break forloop2
                            }
                        }
                    case EIGHT[4]:
                        if len(scanner.Text()[:i]) >= 5 {
                            if scanner.Text()[i-4:i+1] == EIGHT {
                                digit = digit + "8"
                                break forloop2
                            }
                        }
                    }
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

	
}