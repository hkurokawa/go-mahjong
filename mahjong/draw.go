package mahjong

import (
	"errors"
	"fmt"
	"math/rand"
)

// Randomly draw Pais from the pile
func drawPais(pile []Pai, n int, r *rand.Rand) ([]Pai, []Pai, error) {
	if len(pile) < n {
		return nil, nil, errors.New(fmt.Sprintf("No remaining pais for draw %d Pais. Remaining: %d.", n, len(pile)))
	}

	p := make([]Pai, n)
	for i := 0; i < n; i++ {
		// Roll a dice and see which Pai is selected
		j := r.Int31n(int32(len(pile)))
		p[i] = pile[j]
		pile = append(pile[:j], pile[j+1:]...)
	}
	return p, pile, nil
}

// Randomly draw a Pai from the pile.
func drawPai(pile []Pai, r *rand.Rand) (Pai, []Pai, error) {
	p, np, err := drawPais(pile, 1, r)
	if err != nil {
		return Pai{}, nil, err
	}
	if len(p) != 1 {
		return Pai{}, nil, errors.New(fmt.Sprintf("Only one Pai should be picked up: %s.", p))
	}
	return p[0], np, nil
}
