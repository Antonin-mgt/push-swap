package parsing

import (
	"fmt"
	"strconv"
	"strings"
)

// Parser convertit un texte contenant des entiers séparés par des espaces
// en une tranche d'entiers. Elle renvoie une erreur si un élément n'est pas
// un entier ou si un nombre apparaît plusieurs fois.
func Parser(entree string) ([]int, error) {
	morceaux := strings.Fields(entree)
	nombres := make([]int, len(morceaux)) // crée une tranche d'entier de la même taille que morceaux
	for i, morceau := range morceaux {
		nombre, err := strconv.Atoi(morceau) // strconv.Atoi convertit le morceau en entier.
		if err != nil {
			return nil, err
		}
		nombres[i] = nombre
	}
	// La map vus retient les nombres déjà rencontrés pour repérer les doublons.
	vus := make(map[int]bool)
	for _, nombre := range nombres {
		if vus[nombre] {
			return nil, fmt.Errorf("erreur : doublon détecté pour le nombre %d", nombre)
		}
		vus[nombre] = true
	}
	return nombres, nil
}
