package words

import (
	"testing"
)

// Test new words returns words struct
func TestNewWords(t *testing.T) {
	t.Log("Testing NewWords()... returns a new Words struct")
	parser := NewWords()
	if parser == nil {
		t.Errorf("NewWords() returned nil")
	}
}

// Test validate method returns an error if the input is not a number
func TestValidate(t *testing.T) {
	t.Log("Testing Validate()... returns an error if the input is not a number")
	parser := NewWords()
	err := parser.Validate("not a number")
	if err == nil {
		t.Errorf("Validate() did not return an error")
	}
}

// Test validate method returns an error if the input is not between 0-100,000
func TestValidate2(t *testing.T) {
	t.Log("Testing Validate()... returns an error if the input is not between 0-100,000")
	parser := NewWords()
	err := parser.Validate("100001")
	if err == nil {
		t.Errorf("Validate() did not return an error")
	}
}

// Test validate returns no error when given a valid input
func TestValidate3(t *testing.T) {
	t.Log("Testing Validate()... returns no error when given a valid input")
	parser := NewWords()
	err := parser.Validate("100000")
	if err != nil {
		t.Errorf("Validate() returned an error")
	}
}

// Test parse method returns the correct string for 0
func TestParse(t *testing.T) {
	t.Log("Testing Parse()... returns the correct string for 0")
	parser := NewWords()
	parser.Validate("0")
	numberInEnglish := parser.Parse()
	if numberInEnglish != "zero" {
		t.Errorf("Parse() returned %s, expected zero", numberInEnglish)
	}
}

// Test parse method returns the correct string for 1
func TestParse2(t *testing.T) {
	t.Log("Testing Parse()... returns the correct string for 1")
	parser := NewWords()
	parser.Validate("1")
	numberInEnglish := parser.Parse()
	if numberInEnglish != "one" {
		t.Errorf("Parse() returned %s, expected one", numberInEnglish)
	}
}

// Test parse method returns the correct string for 9
func TestParse3(t *testing.T) {
	t.Log("Testing Parse()... returns the correct string for 9")
	parser := NewWords()
	parser.Validate("9")
	numberInEnglish := parser.Parse()
	if numberInEnglish != "nine" {
		t.Errorf("Parse() returned %s, expected nine", numberInEnglish)
	}
}

// Test parse method returns the correct string for 10
func TestParse4(t *testing.T) {
	t.Log("Testing Parse()... returns the correct string for 10")
	parser := NewWords()
	parser.Validate("10")
	numberInEnglish := parser.Parse()
	if numberInEnglish != "ten" {
		t.Errorf("Parse() returned %s, expected ten", numberInEnglish)
	}
}

// Test parse method returns the correct string for 11
func TestParse5(t *testing.T) {
	t.Log("Testing Parse()... returns the correct string for 11")
	parser := NewWords()
	parser.Validate("11")
	numberInEnglish := parser.Parse()
	if numberInEnglish != "eleven" {
		t.Errorf("Parse() returned %s, expected eleven", numberInEnglish)
	}
}

// Test parse method returns the correct string for 19
func TestParse6(t *testing.T) {
	t.Log("Testing Parse()... returns the correct string for 19")
	parser := NewWords()
	parser.Validate("19")
	numberInEnglish := parser.Parse()
	if numberInEnglish != "nineteen" {
		t.Errorf("Parse() returned %s, expected nineteen", numberInEnglish)
	}
}

// Test parse method returns the correct string for 20
func TestParse7(t *testing.T) {
	t.Log("Testing Parse()... returns the correct string for 20")
	parser := NewWords()
	parser.Validate("20")
	numberInEnglish := parser.Parse()
	if numberInEnglish != "twenty" {
		t.Errorf("Parse() returned %s, expected twenty", numberInEnglish)
	}
}

// Test parse method returns the correct string for 101
func TestParse8(t *testing.T) {
	t.Log("Testing Parse()... returns the correct string for 101")
	parser := NewWords()
	parser.Validate("101")
	numberInEnglish := parser.Parse()
	if numberInEnglish != "one hundred and one" {
		t.Errorf("Parse() returned %s, expected one hundred and one", numberInEnglish)
	}
}

// Test parse method returns the correct string for 119
func TestParse9(t *testing.T) {
	t.Log("Testing Parse()... returns the correct string for 119")
	parser := NewWords()
	parser.Validate("119")
	numberInEnglish := parser.Parse()
	if numberInEnglish != "one hundred and nineteen" {
		t.Errorf("Parse() returned %s, expected one hundred and nineteen", numberInEnglish)
	}
}

// Test parse method returns the correct string for 999
func TestParse10(t *testing.T) {
	t.Log("Testing Parse()... returns the correct string for 999")
	parser := NewWords()
	parser.Validate("999")
	numberInEnglish := parser.Parse()
	if numberInEnglish != "nine hundred and ninety-nine" {
		t.Errorf("Parse() returned %s, expected nine hundred and ninety-nine", numberInEnglish)
	}
}

// Test parse method returns the correct string for 1000
func TestParse11(t *testing.T) {
	t.Log("Testing Parse()... returns the correct string for 1000")
	parser := NewWords()
	parser.Validate("1000")
	numberInEnglish := parser.Parse()
	if numberInEnglish != "one thousand" {
		t.Errorf("Parse() returned %s, expected one thousand", numberInEnglish)
	}
}

// Test parse method returns the correct string for 1001
func TestParse12(t *testing.T) {
	t.Log("Testing Parse()... returns the correct string for 1001")
	parser := NewWords()
	parser.Validate("1001")
	numberInEnglish := parser.Parse()
	if numberInEnglish != "one thousand and one" {
		t.Errorf("Parse() returned %s, expected one thousand and one", numberInEnglish)
	}
}

// Test parse method returns the correct string for 1010
func TestParse13(t *testing.T) {
	t.Log("Testing Parse()... returns the correct string for 1010")
	parser := NewWords()
	parser.Validate("1010")
	numberInEnglish := parser.Parse()
	if numberInEnglish != "one thousand and ten" {
		t.Errorf("Parse() returned %s, expected one thousand and ten", numberInEnglish)
	}
}

// Test parse method returns the correct string for 1111
func TestParse14(t *testing.T) {
	t.Log("Testing Parse()... returns the correct string for 1111")
	parser := NewWords()
	parser.Validate("1111")
	numberInEnglish := parser.Parse()
	if numberInEnglish != "one thousand, one hundred and eleven" {
		t.Errorf("Parse() returned %s, expected one thousand one hundred and eleven", numberInEnglish)
	}
}

// Test parse method returns the correct string for 9999
func TestParse15(t *testing.T) {
	t.Log("Testing Parse()... returns the correct string for 9999")
	parser := NewWords()
	parser.Validate("9999")
	numberInEnglish := parser.Parse()
	if numberInEnglish != "nine thousand, nine hundred and ninety-nine" {
		t.Errorf("Parse() returned %s, expected nine thousand, nine hundred and ninety-nine", numberInEnglish)
	}
}

// Test parse method returns the correct string for 10000
func TestParse16(t *testing.T) {
	t.Log("Testing Parse()... returns the correct string for 10000")
	parser := NewWords()
	parser.Validate("10000")
	numberInEnglish := parser.Parse()
	if numberInEnglish != "ten thousand" {
		t.Errorf("Parse() returned %s, expected ten thousand", numberInEnglish)
	}
}

// Test parse method returns the correct string for 10001
func TestParse17(t *testing.T) {
	t.Log("Testing Parse()... returns the correct string for 10001")
	parser := NewWords()
	parser.Validate("10001")
	numberInEnglish := parser.Parse()
	if numberInEnglish != "ten thousand and one" {
		t.Errorf("Parse() returned %s, expected ten thousand and one", numberInEnglish)
	}
}

// Test parse method returns the correct string for 10010
func TestParse18(t *testing.T) {
	t.Log("Testing Parse()... returns the correct string for 10010")
	parser := NewWords()
	parser.Validate("10010")
	numberInEnglish := parser.Parse()
	if numberInEnglish != "ten thousand and ten" {
		t.Errorf("Parse() returned %s, expected ten thousand and ten", numberInEnglish)
	}
}

// Test parse method returns the correct string for 10101
func TestParse19(t *testing.T) {
	t.Log("Testing Parse()... returns the correct string for 10101")
	parser := NewWords()
	parser.Validate("10101")
	numberInEnglish := parser.Parse()
	if numberInEnglish != "ten thousand, one hundred and one" {
		t.Errorf("Parse() returned %s, expected ten thousand, one hundred and one", numberInEnglish)
	}
}

// Test parse method returns the correct string for 11111
func TestParse20(t *testing.T) {
	t.Log("Testing Parse()... returns the correct string for 11111")
	parser := NewWords()
	parser.Validate("11111")
	numberInEnglish := parser.Parse()
	if numberInEnglish != "eleven thousand, one hundred and eleven" {
		t.Errorf("Parse() returned %s, expected eleven thousand, one hundred and eleven", numberInEnglish)
	}
}

// Test parse method returns the correct string for 99999
func TestParse21(t *testing.T) {
	t.Log("Testing Parse()... returns the correct string for 99999")
	parser := NewWords()
	parser.Validate("99999")
	numberInEnglish := parser.Parse()
	if numberInEnglish != "ninety-nine thousand, nine hundred and ninety-nine" {
		t.Errorf("Parse() returned %s, expected ninety-nine thousand, nine hundred and ninety-nine", numberInEnglish)
	}
}

// Test parse method returns the correct string for 100000
func TestParse22(t *testing.T) {
	t.Log("Testing Parse()... returns the correct string for 100000")
	parser := NewWords()
	parser.Validate("100000")
	numberInEnglish := parser.Parse()
	if numberInEnglish != "one hundred thousand" {
		t.Errorf("Parse() returned %s, expected one hundred thousand", numberInEnglish)
	}
}
