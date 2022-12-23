package words

import (
	"fmt"
	"strconv"
)

type Words struct {
	oneToNine        []string
	elevenToNineteen []string
	tenToNinety      []string
	hundred          string
	and              string
	thousand         string
	number           int
	numberOfDigits   int
}

func NewWords() *Words {
	return &Words{
		oneToNine:        []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"},
		elevenToNineteen: []string{"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"},
		tenToNinety:      []string{"ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"},
		hundred:          "hundred",
		and:              "and",
		thousand:         "thousand",
	}
}

// Validates the CLI input to make sure it is a number within the expected range
func (w *Words) Validate(numberAsString string) error {
	number, err := strconv.Atoi(numberAsString)
	if err != nil {
		return fmt.Errorf("Input must be a number")
	}

	if number < 0 || number > 100000 {
		return fmt.Errorf("Input number must be between 1-100,000")
	}
	w.number = number
	w.numberOfDigits = len(numberAsString)
	return nil
}

func (w *Words) Parse() string {
	if w.number == 0 {
		return "zero"
	}

	numberInEnglish := ""
	switch w.numberOfDigits {
	case 1:
		numberInEnglish = w.parseOneDigitNumber(w.number)
	case 2:
		numberInEnglish = w.parseTwoDigitNumber(w.number)
	case 3:
		numberInEnglish = w.parseThreeDigitNumber(w.number)
	case 4:
		numberInEnglish = w.parseFourDigitNumber(w.number)
	case 5:
		numberInEnglish = w.parseFiveDigitNumber(w.number)
	case 6:
		numberInEnglish = w.parseSixDigitNumber(w.number)
	}
	return numberInEnglish
}

// function that takes 1-9 as an integer and returns the number in english
func (w *Words) parseOneDigitNumber(number int) string {
	return w.oneToNine[number-1]
}

// function that takes 10-99 as an integer and returns the number in english
func (w *Words) parseTwoDigitNumber(number int) string {
	if number > 10 && number < 20 {
		return w.elevenToNineteen[number-11]
	}
	if number%10 == 0 {
		return w.tenToNinety[number/10-1]
	}

	return w.tenToNinety[number/10-1] + "-" + w.parseOneDigitNumber(number%10)
}

// function that takes 100-999 as an integer and returns the number in english
func (w *Words) parseThreeDigitNumber(number int) string {
	prefix := w.parseOneDigitNumber(number/100) + " " + w.hundred
	return w.switchForModuloOfEndNumbers(prefix, number%100)
}

// function that takes 1000-9999 as an integer and returns the number in english
func (w *Words) parseFourDigitNumber(number int) string {
	prefix := w.parseOneDigitNumber(number/1000) + " " + w.thousand
	return w.switchForModuloOfEndNumbers(prefix, number%1000)
}

// function that takes 10000-99999 as an integer and returns the number in english
func (w *Words) parseFiveDigitNumber(number int) string {
	prefix := w.parseTwoDigitNumber(number/1000) + " " + w.thousand
	return w.switchForModuloOfEndNumbers(prefix, number%1000)
}

// function parses six digit numbers, only returns 100000 at the moment as that is the limit
func (w *Words) parseSixDigitNumber(number int) string {
	return "one hundred thousand"
}

func (w *Words) switchForModuloOfEndNumbers(prefix string, moduloResult int) string {
	switch m := moduloResult; {
	case m == 0:
		return prefix
	case m < 10:
		return prefix + ", " + w.and + " " + w.parseOneDigitNumber(m)
	case m < 100:
		return prefix + ", " + w.and + " " + w.parseTwoDigitNumber(m)
	default:
		return prefix + ", " + w.and + " " + w.parseThreeDigitNumber(m)
	}
}
