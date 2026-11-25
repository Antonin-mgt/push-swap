package piles

// positionDuMinimum renvoie la position (en partant de 0) du plus petit élément de la pile A.
// Elle renvoie -1 si la pile A est vide.
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

// Trier trie la pile A en deux phases :
//  1. tant que A n'est pas triée, on amène son plus petit élément en haut
//     (avec Ra ou Rra, selon le chemin le plus court) puis on l'envoie sur B ;
//  2. on remet ensuite tous les éléments de B sur A. B s'est remplie du plus
//     petit au plus grand, donc A se reconstruit à l'envers et se retrouve triée.
func (p *Piles) Trier() {
	for !p.EstTriee() {
		pos := p.positionDuMinimum()
		// Si le minimum est dans la première moitié de A, on le remonte avec Ra,
		// sinon il est plus rapide de le ramener par le bas avec Rra.
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
