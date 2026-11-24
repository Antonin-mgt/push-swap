package piles

type Piles struct {
	A          []int
	B          []int
	Operations []string
}

func (p *Piles) Ra() {
	if len(p.A) < 2 {
		return
	}
	p.A = append(p.A[1:], p.A[0])
	p.Operations = append(p.Operations, "ra")
}

func (p *Piles) Rb() {
	if len(p.B) < 2 {
		return
	}
	p.B = append(p.B[1:], p.B[0])
	p.Operations = append(p.Operations, "rb")
}

func (p *Piles) Sa() {
	if len(p.A) < 2 {
		return
	}
	p.A[0], p.A[1] = p.A[1], p.A[0]
	p.Operations = append(p.Operations, "sa")
}

func (p *Piles) Sb() {
	if len(p.B) < 2 {
		return
	}
	p.B[0], p.B[1] = p.B[1], p.B[0]
	p.Operations = append(p.Operations, "sb")
}

func (p *Piles) Pa() {
	if len(p.B) == 0 {
		return
	}
	p.A = append([]int{p.B[0]}, p.A...)
	p.B = p.B[1:]
	p.Operations = append(p.Operations, "pa")
}

func (p *Piles) Pb() {
	if len(p.A) == 0 {
		return
	}
	p.B = append([]int{p.A[0]}, p.B...)
	p.A = p.A[1:]
	p.Operations = append(p.Operations, "pb")
}

func (p *Piles) Rra() {
	if len(p.A) < 2 {
		return
	}
	p.A = append([]int{p.A[len(p.A)-1]}, p.A[:len(p.A)-1]...)
	p.Operations = append(p.Operations, "rra")
}

func (p *Piles) Rrb() {
	if len(p.B) < 2 {
		return
	}
	p.B = append([]int{p.B[len(p.B)-1]}, p.B[:len(p.B)-1]...)
	p.Operations = append(p.Operations, "rrb")
}

func (p *Piles) Ss() {
	if len(p.A) >= 2 {
		p.A[0], p.A[1] = p.A[1], p.A[0]
	}
	if len(p.B) >= 2 {
		p.B[0], p.B[1] = p.B[1], p.B[0]
	}
	p.Operations = append(p.Operations, "ss")
}

func (p *Piles) Rr() {
	if len(p.A) >= 2 {
		p.A = append(p.A[1:], p.A[0])
	}
	if len(p.B) >= 2 {
		p.B = append(p.B[1:], p.B[0])
	}
	p.Operations = append(p.Operations, "rr")
}

func (p *Piles) Rrr() {
	if len(p.A) >= 2 {
		p.A = append([]int{p.A[len(p.A)-1]}, p.A[:len(p.A)-1]...)
	}
	if len(p.B) >= 2 {
		p.B = append([]int{p.B[len(p.B)-1]}, p.B[:len(p.B)-1]...)
	}
	p.Operations = append(p.Operations, "rrr")
}
