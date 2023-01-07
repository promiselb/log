package main

import (
	"log"

	w "github.com/promiselb/log"
)

func main() {
	var l = w.NewLogger(nil, "Alex", log.Ltime)
	dictionary := map[string]string{
		"barrage":   "a concentrated artillery bombardment over a wide area.",
		"debugging": "the process of identifying and removing errors from computer hardware or software.",
	}
	// w.PrintMap(l, "dictionary", dictionary)

	l.PrintValue("dictionary", dictionary)
	l.MakePrinter("main").PrintValue("dictionary", dictionary)
}
