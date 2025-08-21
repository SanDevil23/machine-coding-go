package main

type Player struct {
	name  	string
	symbol 	string
}

func NewPlayer(name string, symbol string) *Player {
	return &Player{
		name:   name,
		symbol: symbol,
	}
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetSymbol() string {
	return p.symbol
}