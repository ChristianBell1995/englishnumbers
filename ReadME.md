# English Numbers
A command-line interface (CLI) tool that takes an integer from 0 to 100,000 and converts it into grammatically correct english words.

## To use binary
Clone this repo and the binary should already be built. If you run `./bin/numbers-to-words <number>` you should see an english word for any integer between 1-100000.

## Setup
1. Install go (`brew install go`)
2. Install the dependencies `go get github.com/ChristianBell1995/englishnumbers`
3. Run the tests. `go test ./words -v`
4. Run the application `go run main.go <number>`
5. To build the binary run `go build -o ./bin/numbers-to-words`
6. Then you should be able to run the binary `./bin/numbers-to-words 1000`
