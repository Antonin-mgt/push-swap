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

func TestEstTriee(t *testing.T) {
	piles := &Piles{A: []int{1, 2, 3}, B: []int{}}
	if !piles.EstTriee() {
		t.Errorf("la pile A devrait être triée")
	}

	piles = &Piles{A: []int{3, 2, 1}, B: []int{}}
	if piles.EstTriee() {
		t.Errorf("la pile A ne devrait pas être triée")
	}

	piles = &Piles{A: []int{}, B: []int{}}
	if !piles.EstTriee() {
		t.Errorf("une pile vide devrait être considérée comme triée")
	}
}

func TestPositionDuMinimum(t *testing.T) {
	piles := &Piles{A: []int{3, 1, 2}}
	minIndex := piles.positionDuMinimum()
	if minIndex != 1 {
		t.Errorf("index du minimum incorrect : obtenu %d, attendu 1", minIndex)
	}

	piles = &Piles{A: []int{1, 2, 3}}
	minIndex = piles.positionDuMinimum()
	if minIndex != 0 {
		t.Errorf("index du minimum incorrect : obtenu %d, attendu 0", minIndex)
	}

	piles = &Piles{A: []int{2, 3, 1}}
	minIndex = piles.positionDuMinimum()
	if minIndex != 2 {
		t.Errorf("index du minimum incorrect : obtenu %d, attendu 2", minIndex)
	}
}

func TestTrier(t *testing.T) {
	piles := &Piles{A: []int{5, 2, 9, 1, 7}, B: []int{}}
	piles.Trier()

	attenduA := []int{1, 2, 5, 7, 9}
	attenduB := []int{}

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
