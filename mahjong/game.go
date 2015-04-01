package mahjong

type Kaze rune

const (
	東風 Kaze = '東'
	南風      = '南'
	西風      = '西'
	北風      = '北'
)

type Order Kaze

// A Command specifies an action a player can take.
type Command int

const (
	Tsumo Command = iota
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
	state State // Public information about the current game.
	pile map[Pai]int // Pais in the pile.
}

func (g *Game) Init() error {
	return nil
}

// Randomly pick-up a pai from the pile.
func (g *Game) pick() Pai {
	return Pai{}
}

// Return available commands for the given player.
func (g Game) Commands(p Player) []Command {
	return []Command{}
}

// Play the specified action on the game. If the action cannot be executed, an error returns.
func (g *Game) Play(a Action) error {
	return nil
}

func (g Game) Status() State {
	return State{}
}

// A State specifies public information of the game,
// such as the number of remaining pais in the pile, who discarded which pais (Ho).
type State struct {
	Junnme  int
	NumPais int          // The number of remaining tsumoable pais.
	Honnba  int          // How many times the renchan repeats.
	Kyotaku int          // Deposit score.
	Players []PlayerInfo // Public information about players.
	Dora    []Pai
}

// A Player specifies private information of a player.
type Player struct {
	PlayerInfo
	Tehai []Pai
}

// A PlayerInfo specifies public information of a player.
type PlayerInfo struct {
	Id   int
	Name string
	Kaze
	Score int
	Order
	Ho   []Sutehai
	Furo []Mentsu
}
