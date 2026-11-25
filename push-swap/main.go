package main

import (
	"fmt"
	"os"
	"push-swap/parsing"
	"push-swap/piles"
)

// main lit les nombres en argument, trie la pile A et affiche
// les instructions utilisées, une par ligne. Une entrée invalide affiche Error.
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
	// On affiche chaque instruction enregistrée pendant le tri.
	for _, instruction := range p.Operations {
		fmt.Println(instruction)
	}
}
