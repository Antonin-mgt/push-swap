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
