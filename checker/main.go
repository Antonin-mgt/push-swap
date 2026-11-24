package main

import (
	"bufio"
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

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ligne := scanner.Text()
		if ligne == "" {
			continue
		}
		if !executer(p, ligne) {
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}
	if p.EstTriee() && len(p.B) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

func executer(p *piles.Piles, instruction string) bool {
	switch instruction {
	case "sa":
		p.Sa()
	case "sb":
		p.Sb()
	case "ss":
		p.Ss()
	case "pa":
		p.Pa()
	case "pb":
		p.Pb()
	case "ra":
		p.Ra()
	case "rb":
		p.Rb()
	case "rr":
		p.Rr()
	case "rra":
		p.Rra()
	case "rrb":
		p.Rrb()
	case "rrr":
		p.Rrr()
	default:
		return false
	}
	return true
}
