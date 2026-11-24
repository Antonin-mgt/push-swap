package piles

import (
	"testing"
)

func TestRrr(t *testing.T) {
	piles := &Piles{A: []int{1, 2, 3}, B: []int{4, 5, 6}}
	piles.Rrr()

	attenduA := []int{3, 1, 2}
	attenduB := []int{6, 4, 5}

	if len(piles.A) != len(attenduA) {
		t.Errorf("taille incorrecte : obtenu %v, attendu %v", piles.A, attenduA)
		return
	}
	for i := range attenduA {
		if piles.A[i] != attenduA[i] {
			t.Errorf("case %d incorrecte : obtenu %v, attendu %v", i, piles.A, attenduA)
		}
	}
	if len(piles.B) != len(attenduB) {
		t.Errorf("taille incorrecte : obtenu %v, attendu %v", piles.B, attenduB)
		return
	}
	for i := range attenduB {
		if piles.B[i] != attenduB[i] {
			t.Errorf("case %d incorrecte : obtenu %v, attendu %v", i, piles.B, attenduB)
		}
	}
}

func TestRr(t *testing.T) {
	piles := &Piles{A: []int{1, 2, 3}, B: []int{4, 5, 6}}
	piles.Rr()

	attenduA := []int{2, 3, 1}
	attenduB := []int{5, 6, 4}

	if len(piles.A) != len(attenduA) {
		t.Errorf("taille incorrecte : obtenu %v, attendu %v", piles.A, attenduA)
		return
	}
	for i := range attenduA {
		if piles.A[i] != attenduA[i] {
			t.Errorf("case %d incorrecte : obtenu %v, attendu %v", i, piles.A, attenduA)
		}
	}
	if len(piles.B) != len(attenduB) {
		t.Errorf("taille incorrecte : obtenu %v, attendu %v", piles.B, attenduB)
		return
	}
	for i := range attenduB {
		if piles.B[i] != attenduB[i] {
			t.Errorf("case %d incorrecte : obtenu %v, attendu %v", i, piles.B, attenduB)
		}
	}
}

func TestSs(t *testing.T) {
	piles := &Piles{A: []int{1, 2, 3}, B: []int{4, 5, 6}}
	piles.Ss()

	attenduA := []int{2, 1, 3}
	attenduB := []int{5, 4, 6}

	if len(piles.A) != len(attenduA) {
		t.Errorf("taille incorrecte : obtenu %v, attendu %v", piles.A, attenduA)
		return
	}
	for i := range attenduA {
		if piles.A[i] != attenduA[i] {
			t.Errorf("case %d incorrecte : obtenu %v, attendu %v", i, piles.A, attenduA)
		}
	}
	if len(piles.B) != len(attenduB) {
		t.Errorf("taille incorrecte : obtenu %v, attendu %v", piles.B, attenduB)
		return
	}
	for i := range attenduB {
		if piles.B[i] != attenduB[i] {
			t.Errorf("case %d incorrecte : obtenu %v, attendu %v", i, piles.B, attenduB)
		}
	}
}
