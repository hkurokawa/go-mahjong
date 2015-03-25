package mahjong

type Type int

const (
	Shuntsu = iota
	Kotsu
	Kantsu
	Toitsu
)

type Suite int

const (
	萬子 = iota
	索子
	筒子
	字子
)

type Rank int

const (
	東 = 10 + iota
	南
	西
	北
	白
	発
	中
)

type Pai struct {
	Suite
	Rank int
}

type Mentsu struct {
	Type
	Member []Pai
}

type Hand struct {
	Expose []Mentsu
	Tehai  []Pai
}
