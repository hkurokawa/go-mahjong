package mahjong
import (
    "testing"
    "math/rand"
    "time"
)

func TestDrawPaisAll(t *testing.T) {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    pile := make(map[Pai]int)
    pile[Pai{Zizu, Tong}] = 1
    pile[Pai{Zizu, Nang}] = 1
    pile[Pai{Zizu, Sha}] = 1
    pile[Pai{Zizu, Pei}] = 1
    pile[Pai{Zizu, Haku}] = 1
    pile[Pai{Zizu, Fa}] = 1
    pile[Pai{Zizu, Chung}] = 1
    pile[Pai{Manzu, 1}] = 1
    pile[Pai{Manzu, 9}] = 1
    pile[Pai{Sozu, 1}] = 1
    pile[Pai{Sozu, 9}] = 1
    pile[Pai{Pinzu, 1}] = 1
    pile[Pai{Pinzu, 9}] = 1
    expected := make(map[Pai]int)
    for k, v := range pile {
        expected[k] = v
    }
    hand, err := drawPais(pile, 13, r)
    if err != nil {
        t.Fatalf("Failed to draw %d Pais: %s.", 13, err)
    }
    testHand(t, hand, expected)
    num := 0
    for _, v := range pile {
        num += v
    }
    if num != 0 {
        t.Errorf("Pile should be empty: %s.", pile)
    }
}

func testHand(t *testing.T, a map[Pai]int, e map[Pai]int) {
    for k, v := range a {
        if v != e[k] {
            t.Fatalf("Actual value differs at [%s]: %d ≠ %d.", k, v, e[k])
        }
        delete(e, k)
    }
    // Verify all the Pais in the expected is removed.
    if len(e) != 0 {
        t.Fatalf("Remaining expected values: %s.", e)
    }
}