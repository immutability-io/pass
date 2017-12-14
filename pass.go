package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/sethvargo/go-diceware/diceware"
)

func main() {

	words := flag.Int("words", 6, "Number of words in phrase.")
	separator := flag.String("separator", " ", "String to separate words.")
	flag.Parse()
	list, _ := diceware.Generate(*words)
	passphrase := strings.Join(list, *separator)
	fmt.Println(passphrase)
}
