package piles

// Piles contient les piles A et B (le sommet est la case 0) et la liste des
// instructions déjà exécutées.
type Piles struct {
	A          []int
	B          []int
	Operations []string
}

// Ra fait passer le premier élément de la pile A à la fin,
// et ne fait rien si la pile a moins de 2 éléments.
func (p *Piles) Ra() {
	if len(p.A) < 2 {
		return
	}
	p.A = append(p.A[1:], p.A[0])
	p.Operations = append(p.Operations, "ra")
}

// Rb fait passer le premier élément de la pile B à la fin,
// et ne fait rien si la pile a moins de 2 éléments.
func (p *Piles) Rb() {
	if len(p.B) < 2 {
		return
	}
	p.B = append(p.B[1:], p.B[0])
	p.Operations = append(p.Operations, "rb")
}

// Sa échange les deux premiers éléments de A.
// Elle ne fait rien si A contient moins de 2 éléments.
func (p *Piles) Sa() {
	if len(p.A) < 2 {
		return
	}
	p.A[0], p.A[1] = p.A[1], p.A[0]
	p.Operations = append(p.Operations, "sa")
}

// Sb échange les deux premiers éléments de B.
// Elle ne fait rien si B contient moins de 2 éléments.
func (p *Piles) Sb() {
	if len(p.B) < 2 {
		return
	}
	p.B[0], p.B[1] = p.B[1], p.B[0]
	p.Operations = append(p.Operations, "sb")
}

// Pa prend le premier élément de B et le pose sur A.
// Elle ne fait rien si B est vide.
func (p *Piles) Pa() {
	if len(p.B) == 0 {
		return
	}
	p.A = append([]int{p.B[0]}, p.A...)
	p.B = p.B[1:]
	p.Operations = append(p.Operations, "pa")
}

// Pb prend le premier élément de A et le pose sur B.
// Elle ne fait rien si A est vide.
func (p *Piles) Pb() {
	if len(p.A) == 0 {
		return
	}
	p.B = append([]int{p.A[0]}, p.B...)
	p.A = p.A[1:]
	p.Operations = append(p.Operations, "pb")
}

// Rra ramène le dernier élément de A au début.
func (p *Piles) Rra() {
	if len(p.A) < 2 {
		return
	}
	p.A = append([]int{p.A[len(p.A)-1]}, p.A[:len(p.A)-1]...)
	p.Operations = append(p.Operations, "rra")
}

// Rrb ramène le dernier élément de B au début.
func (p *Piles) Rrb() {
	if len(p.B) < 2 {
		return
	}
	p.B = append([]int{p.B[len(p.B)-1]}, p.B[:len(p.B)-1]...)
	p.Operations = append(p.Operations, "rrb")
}

// Ss échange les deux premiers éléments de A et de B en une seule instruction.
// Une pile de moins de 2 éléments reste inchangée.
func (p *Piles) Ss() {
	if len(p.A) >= 2 {
		p.A[0], p.A[1] = p.A[1], p.A[0]
	}
	if len(p.B) >= 2 {
		p.B[0], p.B[1] = p.B[1], p.B[0]
	}
	p.Operations = append(p.Operations, "ss")
}

// Rr déplace le premier élément de A et de B à la fin de leur pile
// respective au même moment en une opération.
func (p *Piles) Rr() {
	if len(p.A) >= 2 {
		p.A = append(p.A[1:], p.A[0])
	}
	if len(p.B) >= 2 {
		p.B = append(p.B[1:], p.B[0])
	}
	p.Operations = append(p.Operations, "rr")
}

// Rrr déplace le dernier élément de A et de B au début de leur pile
// respective au même moment en une opération.
func (p *Piles) Rrr() {
	if len(p.A) >= 2 {
		p.A = append([]int{p.A[len(p.A)-1]}, p.A[:len(p.A)-1]...)
	}
	if len(p.B) >= 2 {
		p.B = append([]int{p.B[len(p.B)-1]}, p.B[:len(p.B)-1]...)
	}
	p.Operations = append(p.Operations, "rrr")
}

// EstTriee indique si la pile A est triée dans l'ordre croissant
// (B n'est pas regardée).
func (p *Piles) EstTriee() bool {
	for i := 0; i < len(p.A)-1; i++ {
		if p.A[i] > p.A[i+1] {
			return false
		}
	}
	return true
}
