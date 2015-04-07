package mahjong

import (
	"math/rand"
	"time"
)

type Kaze rune

const (
	TongPu Kaze = '東'
	NangPu      = '南'
	ShaPu       = '西'
	PeiPu       = '北'
)

type Order Kaze

// A Command specifies an action a player can take.
type Command struct {
	Type CommandType
	With []Pai
}
type CommandType int

const (
	Tsumo CommandType = iota
	TsumoHoura
	RonHoura
	Chi
	Pong
	AnngKan
	MingKan
	Tahai
	TahaiReach
	Pass
)

// An Action specifies who does what.
type Action struct {
	Player  Player
	Command Command
}

// A Game specifies public and private information about the current game (Hanchang),
// such as players, pais in the pile, discarded piles (Ho) information.
// They are changed when a player does an action.
type Game struct {
	state State       // Public information about the current game.
	pile  map[Pai]int // Pais in the pile.
	r     *rand.Rand
}

func (g *Game) Init() error {
	// Create a Random
	g.r = rand.New(rand.NewSource(time.Now().UnixNano()))
	// Prepare pile
	numPais := 0
	for _, s := range []Suite{Manzu, Sozu, Pinzu} {
		for i := 0; i < 10; i++ {
			g.pile[Pai{s, Rank(i)}] = 4
			numPais += 4
		}
	}
	for _, r := range []Rank{Tong, Nang, Sha, Pei, Haku, Fa, Chung} {
		g.pile[Pai{Zizu, r}] = 4
		numPais += 4
	}
	// Create default players
	players := []Player{}
	kz := []Kaze{TongPu, NangPu, ShaPu, PeiPu}
	for i, n := range []string{"Alice", "Bob", "Carol", "Ted"} {
		hand, err := drawPais(g.pile, 13, g.r)
		if err != nil {
			// FIXME
			panic(err)
		}
		players[i] = Player{
			PlayerInfo: PlayerInfo{
				Id:    i,
				Name:  n,
				Kaze:  kz[i],
				Score: 25000,
				Order: Order(kz[i]),
				Ho:    []Sutehai{},
				Furo:  []Mentsu{},
			},
			Tehai: hand,
		}
	}

	// Reset game status
	d, err := g.draw()
	if err != nil {
		// FIXME
		panic(err)
	}
	g.state = State{
		Junnme:  1,
		NumPais: numPais,
		Honnba:  0,
		Kyotaku: 0,
		Dora:    []Pai{d},
	}

	return nil
}

// Randomly pick-up a pai from the pile.
func (g *Game) draw() (Pai, error) {
	return drawPai(g.pile, g.r)
}

// Return available commands for the given player.
func (g Game) Commands(p Player) []Command {
	// Always add Pass action
	cmds := []Command{Command{Pass, nil}}
	// See if the last action is made by the previous player or the others.
	// FIXME
	next := true
	// See the last discarded Pai.
	// FIXME
	pai := Pai{}
	// If able to Chi
	if next {
		cmds = append(cmds, p.Tehai.Chiable(pai)...)
	}
	// If able to Pong
	cmds = append(cmds, p.Tehai.Ponnable(pai)...)
	// If able to Kan
	cmds = append(cmds, p.Tehai.Kannable(pai)...)
	// If able to Ron
	if p.Ronnable(pai) {
		cmds = append(cmds, Command{RonHoura, nil})
	}


	return cmds
}

// Play the specified action on the game. If the action cannot be executed, an error returns.
func (g *Game) Play(a Action) error {
	// FIXME
	return nil
}

func (g Game) Status() State {
	// FIXME
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

// A Tehai specifies a sequence of Pais.
type Tehai []Pai

func (t Tehai) Chiable(p Pai) []Command {
	//FIXME
	return nil
}

func (t Tehai) Ponnable(p Pai) []Command {
	//FIXME
	return nil
}

func (t Tehai) Kannable(p Pai) []Command {
	//FIXME
	return nil
}

// A Player specifies private information of a player.
type Player struct {
	PlayerInfo
	Tehai Tehai
}

func (p Player) Ronnable(pai Pai) bool {
	//FIXME
	return false
}

// A PlayerInfo specifies public information of a player.
type PlayerInfo struct {
	Id    int
	Name  string
	Kaze  Kaze
	Score int
	Order Order
	Ho    []Sutehai
	Furo  []Mentsu
}
