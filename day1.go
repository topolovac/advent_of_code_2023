package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fb, err := os.ReadFile("./data/day1_input.txt")
	if err != nil {
		log.Panic(err.Error())
	}

	file_data := string(fb)
	lines := strings.Split(file_data, "\n")
	sum := 0
	for _, value := range lines {
		digit_string := ""
		for _, rune := range value {
			if unicode.IsDigit(rune) {
				digit_string += string(rune)
			}
		}

		fl := digit_string[:1] + digit_string[len(digit_string)-1:]
		digit, err := strconv.ParseInt(fl, 10, 64)
		if err != nil {
			log.Fatal(err.Error())
		}
        fmt.Printf("%s -> %s -> %s\n", value, digit_string, fl)
		sum += int(digit)
	}

	fmt.Println("Total sum is: ", sum)
}
