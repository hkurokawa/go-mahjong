package mahjong

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
)

// Randomly draw Pais from the pile
func drawPais(pile map[Pai]int, n int, r *rand.Rand) (map[Pai]int, error) {
	// Build two arrays representing a number line divided into some fragments.
	// Each fragment represent a Pai of which length is proportional to the remaining number of the Pai.
	kls := []Pai{}
	vls := []int{}
	var num int = 0
	for k, v := range pile {
		if v > 0 {
			num += v
			kls = append(kls, k)
			vls = append(vls, num)
		}
	}
	if num < n {
		return nil, errors.New(fmt.Sprintf("No remaining pais for draw %d Pais. Remaining: %d.", n, num))
	}

	p := make(map[Pai]int)
	for ; n > 0; n-- {
		// Roll a dice and see which fragment contains the value in its range.
		rn := r.Int31n(int32(num))
		i := sort.Search(len(vls), func(i int) bool { return vls[i] > int(rn) })
		p[kls[i]]++
		pile[kls[i]]--
		// Decrement the remaining number of Pais following the selected Pai.
		for ; i < len(vls); i++ {
			if vls[i] > 0 {
				vls[i]--
			}
		}
		num--
	}
	return p, nil
}

// Randomly draw a Pai from the pile.
func drawPai(pile map[Pai]int, r *rand.Rand) (Pai, error) {
	m, err := drawPais(pile, 1, r)
	if err != nil {
		return Pai{}, err
	}
    if len(m) != 1 {
        return Pai{}, errors.New(fmt.Sprintf("Only one Pai should be picked up: %s.", m))
    }
	for k := range m {
		return k, nil
	}
	return Pai{}, errors.New("No Pai is drawed.")
}
