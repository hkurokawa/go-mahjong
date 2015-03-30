package mahjong

type Type int

const (
	Shuntsu Type = iota
	Kotsu
	Kantsu
	Toitsu
)

type Suite int

const (
	萬子 Suite = iota
	索子
	筒子
	字子
)

type Rank int

const (
	一 Rank = iota
	二
	三
	四
	五
	六
	七
	八
	九
	東
	南
	西
	北
	白
	発
	中
)

func (r Rank) isKaze(k Kaze) bool {
	switch k {
	case 東風:
		return r == 東
	case 南風:
		return r == 南
	case 西風:
		return r == 西
	case 北風:
		return r == 北
	default:
		return false
	}
}

type Pai struct {
	Suite
	Rank int
}

type Mentsu struct {
	Type
	Member []Pai
}

type Sutehai struct {
	Pai
	Tedashi bool
}
