package mahjong

type Kaze int

const (
	東風 = 東
	南風 = 南
	西風 = 西
	北風 = 北
)

type Order Kaze

// A Command specifies an action a player can take.
type Command int

const (
	Tsumo = iota
	TsumoHoura
	RonHoura
	Chi
	Pong
	AnngKan
	MingKan
	Tahai
	TahaiReach
)

// An Action specifies who does what.
type Action struct {
	Player
	Command
}

// A Game specifies public and private information about the current game (Hanchang),
// such as players, pais in the pile, discarded piles (Ho) information.
// They are changed when a player does an action.
type Game struct {
}

// A State specifies public information of the game,
// such as the number of remaining pais in the pile, who discarded which pais (Ho).
type State struct {
}

// A Player specifies private information of a player.
type Player struct {
	Id   int
	Name string
	Kaze
	Score int
	Order
	Hand
}

// A PlayerInfo specifies public information of a player.
type PlayerInfo struct {
}

func (g Game) Commands() []Command {
	return []Command{}
}

func (g *Game) Play() error {
	return nil
}

func (g Game) Status() State {
	return State{}
}
