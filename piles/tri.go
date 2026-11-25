package piles

func (p *Piles) positionDuMinimum() int {
	if len(p.A) == 0 {
		return -1
	}
	minIndex := 0
	for i := 1; i < len(p.A); i++ {
		if p.A[i] < p.A[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}

func (p *Piles) Trier() {
	for !p.EstTriee() {
		pos := p.positionDuMinimum()
		if pos <= len(p.A)/2 {
			for i := 0; i < pos; i++ {
				p.Ra()
			}
		} else {
			for i := 0; i < len(p.A)-pos; i++ {
				p.Rra()
			}
		}
		p.Pb()
	}
	for len(p.B) > 0 {
		p.Pa()
	}
}
