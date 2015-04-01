package mahjong
import (
    "testing"
    "math/rand"
    "time"
)

func TestDrawPaisAll(t *testing.T) {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    pile := make(map[Pai]int)
    pile[Pai{字子, 東}] = 1
    pile[Pai{字子, 南}] = 1
    pile[Pai{字子, 西}] = 1
    pile[Pai{字子, 北}] = 1
    pile[Pai{字子, 白}] = 1
    pile[Pai{字子, 発}] = 1
    pile[Pai{字子, 中}] = 1
    pile[Pai{萬子, 一}] = 1
    pile[Pai{萬子, 九}] = 1
    pile[Pai{索子, 一}] = 1
    pile[Pai{索子, 九}] = 1
    pile[Pai{筒子, 一}] = 1
    pile[Pai{筒子, 九}] = 1
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