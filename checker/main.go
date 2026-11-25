package main

import (
	"bufio"
	"fmt"
	"os"
	"push-swap/parsing"
	"push-swap/piles"
)

// main lit les nombres de la pile A en argument et les instructions sur
// l'entrée standard, les exécute, puis affiche OK si A est triée et B vide,
// KO sinon. Une erreur (nombre invalide, doublon, instruction inconnue) affiche Error.
func main() {
	if len(os.Args) < 2 {
		return
	}
	// os.Args[1] est l'argument donné au programme (la liste de nombres).
	nombres, err := parsing.Parser(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}
	p := &piles.Piles{A: nombres}
	// bufio.Scanner lit l'entrée standard ligne par ligne :
	// chaque ligne est une instruction envoyée par push-swap.
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ligne := scanner.Text()
		// On ignore les lignes vides.
		if ligne == "" {
			continue
		}
		if !executer(p, ligne) {
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
	}
	// On vérifie que la lecture de l'entrée s'est bien passée.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}
	// Le résultat est OK seulement si A est triée ET que B est vide.
	// Si des éléments restent sur B, la pile n'est pas triée.
	if p.EstTriee() && len(p.B) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

// executer lance la méthode de Piles qui correspond à l'instruction
// et renvoie true. Elle renvoie false si l'instruction est inconnue.
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
