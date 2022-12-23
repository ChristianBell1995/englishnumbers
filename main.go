package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ChristianBell1995/englishnumbers/words"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "Numbers in English",
		Usage:  "Takes in an integer between 0-100,000 and returns the number in English",
		Action: handler,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func handler(cCtx *cli.Context) error {
	firstArg := cCtx.Args().Get(0)
	parser := words.NewWords()

	if err := parser.Validate(firstArg); err != nil {
		return cli.Exit(err, 1)
	}
	numberInEnglish := parser.Parse()

	fmt.Println(numberInEnglish)

	return nil
}
