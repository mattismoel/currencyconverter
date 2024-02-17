package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mattismoel/currencyconverter/model"
)

var usageString string = `
usage: currencyconverter <amount> <from> <to> [flags]

example:
$ currencyconverter 100 USD EUR

flags:
-h --help         Help
`

func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("not a valid amount of arguments")
		fmt.Println(usageString)
		os.Exit(1)
	}

  amt, err := strconv.ParseFloat(args[0], 32)
  if err != nil {
    panic(err)
  }

  from := args[1]
  to := args[2]

  p := model.NewPrompt(float32(amt), from, to)

  rate, err := p.Rate()
  if err != nil {
    panic(err)
  }

  fmt.Println(rate)

}

