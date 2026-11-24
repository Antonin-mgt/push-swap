package parsing

import "testing"

func TestCasValide(t *testing.T) {
	nombres, err := Parser("2 1 3")
	if err != nil {
		t.Errorf("erreur inattendue : %v", err)
		return
	}

	attendu := []int{2, 1, 3}

	if len(nombres) != len(attendu) {
		t.Errorf("taille incorrecte : obtenu %v, attendu %v", nombres, attendu)
		return
	}
	for i := range attendu {
		if nombres[i] != attendu[i] {
			t.Errorf("case %d incorrecte : obtenu %v, attendu %v", i, nombres, attendu)
		}
	}
}

func TestEntreeVide(t *testing.T) {
	nombres, err := Parser("")
	if err != nil {
		t.Errorf("erreur inattendue : %v", err)
	}
	if len(nombres) != 0 {
		t.Errorf("une tranche vide était attendue, obtenu %v", nombres)
	}
}

func TestMotInvalide(t *testing.T) {
	_, err := Parser("0 one 2 3")
	if err == nil {
		t.Errorf("une erreur était attendue")
	}
}

func TestDoublon(t *testing.T) {
	_, err := Parser("1 2 2")
	if err == nil {
		t.Errorf("une erreur était attendue pour les doublons")
	}
	_, err = Parser("3 1 2 3")
	if err == nil {
		t.Errorf("une erreur était attendue pour les doublons")
	}
}
