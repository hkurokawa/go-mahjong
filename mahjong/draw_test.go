package mahjong

import (
	"math/rand"
	"testing"
)

func TestDrawPaisAll(t *testing.T) {
	r := rand.New(rand.NewSource(99))
	pile := []Pai{Pai{Zizu, Tong}, Pai{Zizu, Nang}, Pai{Zizu, Sha}, Pai{Zizu, Pei}, Pai{Zizu, Haku}, Pai{Zizu, Fa}, Pai{Zizu, Chung}, Pai{Manzu, 1}, Pai{Manzu, 9}, Pai{Sozu, 1}, Pai{Sozu, 9}, Pai{Pinzu, 1}, Pai{Pinzu, 9}}
	expected := []Pai{Pai{Pinzu, 1}, Pai{Manzu, 1}, Pai{Zizu, Sha}, Pai{Zizu, Pei}, Pai{Pinzu, 9}, Pai{Zizu, Fa}, Pai{Manzu, 9}, Pai{Sozu, 1}, Pai{Sozu, 9}, Pai{Zizu, Tong}, Pai{Zizu, Nang}, Pai{Zizu, Chung}, Pai{Zizu, Haku}}
	hand, pile, err := drawPais(pile, 13, r)
	if err != nil {
		t.Fatalf("Failed to draw %d Pais: %s.", 13, err)
	}
	if len(hand) != len(expected) {
		t.Errorf("%s ≠ %s.", hand, expected)
	}
	if len(pile) != 0 {
		t.Errorf("Pile should be empty: %s.", pile)
	}
	for i, v := range hand {
		if v != expected[i] {
			t.Errorf("Hand differs from the expected. %s ≠ %s at %d.", v, expected[i], i)
		}
	}
}
