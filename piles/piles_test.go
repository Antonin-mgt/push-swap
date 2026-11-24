package piles

import (
	"testing"
)

func TestRaRra(t *testing.T) {
	piles := &Piles{A: []int{1, 2, 3}, B: []int{}}
	piles.Ra()
	piles.Rra()

	attendu := []int{1, 2, 3}

	if len(piles.A) != len(attendu) {
		t.Errorf("taille incorrecte : obtenu %v, attendu %v", piles.A, attendu)
		return
	}
	for i := range attendu {
		if piles.A[i] != attendu[i] {
			t.Errorf("case %d incorrecte : obtenu %v, attendu %v", i, piles.A, attendu)
		}
	}
}
