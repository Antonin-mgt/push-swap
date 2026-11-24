package main

import (
	"fmt"
	"os"
	"push-swap/parsing"
	"push-swap/piles"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	nombres, err := parsing.Parser(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}
	p := &piles.Piles{A: nombres}
	p.Trier()
	for _, instruction := range p.Operations {
		fmt.Println(instruction)
	}
}
