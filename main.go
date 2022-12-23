package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "boom",
		Usage:  "make an explosive entrance",
		Action: myFunction,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func myFunction(cCtx *cli.Context) error {
	firstArg := cCtx.Args().Get(0)
	number, err := strconv.Atoi(firstArg)
	if err != nil {
		return cli.Exit("Input must be a number", 1)
	}

	if number < 0 || number > 100000 {
		return cli.Exit("Input number must be between 1-100,000", 1)
	}

	numberInEnglish := ""

	if number == 0 {
		numberInEnglish = "zero"
		fmt.Println(numberInEnglish)
		return nil
	}

	switch len(firstArg) {
	case 1:
		numberInEnglish = parseOneDigitNumber(number)
	case 2:
		numberInEnglish = parseTwoDigitNumber(number)
	case 3:
		numberInEnglish = parseThreeDigitNumber(number)
	case 4:
		numberInEnglish = parseFourDigitNumber(number)
	case 5:
		numberInEnglish = parseFiveDigitNumber(number)
	case 6:
		numberInEnglish = "one hundred thousand"
	}
	fmt.Println(numberInEnglish)

	return nil
}

// function that takes 1-9 as an integer and returns the number in english
func parseOneDigitNumber(number int) string {
	oneToNine := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	return oneToNine[number-1]
}

// function that takes 10-99 as an integer and returns the number in english
func parseTwoDigitNumber(number int) string {
	elevenToNineteen := []string{"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tenToNinety := []string{"ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

	if number > 10 && number < 20 {
		return elevenToNineteen[number-11]
	}
	if number%10 == 0 {
		return tenToNinety[number/10-1]
	}

	return tenToNinety[number/10-1] + "-" + parseOneDigitNumber(number%10)
}

// function that takes 100-999 as an integer and returns the number in english
func parseThreeDigitNumber(number int) string {
	hundred := "hundred"
	and := "and"
	if number%100 == 0 {
		return parseOneDigitNumber(number/100) + " " + hundred
	}
	return parseOneDigitNumber(number/100) + " " + hundred + " " + and + " " + parseTwoDigitNumber(number%100)
}

// function that takes 1000-9999 as an integer and returns the number in english
func parseFourDigitNumber(number int) string {
	thousand := "thousand"
	if number%1000 == 0 {
		return parseOneDigitNumber(number/1000) + " " + thousand
	}
	return parseOneDigitNumber(number/1000) + " " + thousand + " " + parseThreeDigitNumber(number%1000)
}

// function that takes 10000-99999 as an integer and returns the number in english
func parseFiveDigitNumber(number int) string {
	if number%1000 == 0 {
		return parseTwoDigitNumber(number/1000) + " thousand"
	}
	return parseTwoDigitNumber(number/1000) + " thousand " + parseThreeDigitNumber(number%1000)
}
