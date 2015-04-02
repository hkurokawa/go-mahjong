package mahjong

type Type int

const (
	Shuntsu Type = 1 + iota
	Kotsu
	Kantsu
	Toitsu
)

type Suite int

const (
	Manzu Suite = 1 + iota
	Sozu
	Pinzu
	Zizu
)

type Rank int

const (
	Tong Rank = 10 + iota
	Nang
	Sha
	Pei
	Haku
	Fa
	Chung
)

func (r Rank) isKaze(k Kaze) bool {
	switch k {
	case TongPu:
		return r == Tong
	case NangPu:
		return r == Nang
	case ShaPu:
		return r == Sha
	case PeiPu:
		return r == Pei
	default:
		return false
	}
}

type Pai struct {
	Suite Suite
	Rank Rank
}

type Mentsu struct {
	Type Type
	Member []Pai
}

type Sutehai struct {
	Pai
	Tedashi bool
}
