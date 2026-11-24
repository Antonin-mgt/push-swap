package parsing

import (
	"fmt"
	"strconv"
	"strings"
)

func Parser(entree string) ([]int, error) {
	morceaux := strings.Fields(entree)
	nombres := make([]int, len(morceaux))
	for i, morceau := range morceaux {
		nombre, err := strconv.Atoi(morceau)
		if err != nil {
			return nil, err
		}
		nombres[i] = nombre
	}
	vus := make(map[int]bool)
	for _, nombre := range nombres {
		if vus[nombre] {
			return nil, fmt.Errorf("erreur : doublon détecté pour le nombre %d", nombre)
		}
		vus[nombre] = true
	}
	return nombres, nil
}
