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
	What CommandType
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
	state State // Public information about the current game.
	pile  []Pai // Pais in the pile.
	r     *rand.Rand
}

func (g *Game) Init() error {
	// Create a Random
	g.r = rand.New(rand.NewSource(time.Now().UnixNano()))
	// Prepare pile
	g.pile = make([]Pai, 0, 136)
	numPais := 0
	for _, s := range []Suite{Manzu, Sozu, Pinzu} {
		for i := 0; i < 10; i++ {
			for j := 0; j < 4; j++ {
				g.pile = append(g.pile, Pai{s, Rank(i)})
			}
			numPais += 4
		}
	}
	for _, r := range []Rank{Tong, Nang, Sha, Pei, Haku, Fa, Chung} {
		for j := 0; j < 4; j++ {
			g.pile = append(g.pile, Pai{Zizu, r})
		}
		numPais += 4
	}
	// Create default players
	players := []Player{}
	kz := []Kaze{TongPu, NangPu, ShaPu, PeiPu}
	for i, n := range []string{"Alice", "Bob", "Carol", "Ted"} {
		hand, np, err := drawPais(g.pile, 13, g.r)
		if err != nil {
			// FIXME
			panic(err)
		}
		g.pile = np
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
	p, np, err := drawPai(g.pile, g.r)
	if err != nil {
		return Pai{}, err
	}
	g.pile = np
	return p, err
}

// Return available commands for the given player.
func (g Game) Commands(p Player) []Command {
	// If someone has just done an action and no one is in turn
	if g.state.Teban == 0 {
		cmds := []Command{}
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

		// Add Pass action if there is at least one action
		if len(cmds) > 0 {
			cmds = append(cmds, Command{Pass, nil})
		}
		return cmds
	}

	// Otherwise, if it is not the player's turn just return.
	if g.state.Teban != p.Id {
		return []Command{}
	}

	// Your turn!
	// FIXME Should use the last Pai the player drawed.
	pai := Pai{}
	cmds := []Command{}
	// You can discard any Pai as you like.
	for _, v := range p.Tehai {
		cmds = append(cmds, Command{Tahai, []Pai{v}})
	}
	// If your hand is already in Agari, you can declare that.
	if p.Tehai.Hourable() {
		cmds = append(cmds, Command{TsumoHoura, []Pai{pai}})
	}
	if p.Furo == nil || len(p.Furo) == 0 {
		// If there is no Huro and your tehai is Tennpaiable, you can do Reach.
		for _, v := range p.Tehai.Tennpaiable() {
			cmds = append(cmds, Command{TahaiReach, []Pai{v}})
		}
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
	Junnme int
	// The number of remaining tsumoable pais.
	NumPais int
	// How many times the current dealer continues the dealer.
	Honnba int
	// Deposit score.
	Kyotaku int
	// The ID of the player who is in turn.
	// If the value is 0, it means no one is in turn (someone has just done an action).
	Teban int
	// Public information about players.
	Players []PlayerInfo
	// The pai drawn to specify a Dora (Often, the actual Dora is the next to the drawn Pai.)
	Dora []Pai
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

func (t Tehai) Hourable() bool {
	// FIXME
	return false
}

func (t Tehai) Tennpaiable() []Pai {
	// FIXME
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
	// The ID of the player. Must be positive (>0).
	Id    int
	Name  string
	Kaze  Kaze
	Score int
	Order Order
	Ho    []Sutehai
	Furo  []Mentsu
}
