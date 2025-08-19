package main

type Player struct {
	name  	string
	symbol 	string
}

func newPlayer(name string, symbol string) *Player {
	return &Player{
		name:   name,
		symbol: symbol,
	}
}

func (p *Player) getName() string {
	return p.name
}

func (p *Player) getSymbol() string {
	return p.symbol
}