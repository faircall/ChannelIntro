package main

type GameStateT int

const (
	OpeningT GameStateT = iota // this is 0
	NumGameStatesT
)

type GameState struct {
	CurrentState GameStateT
	CurrentTime  float64
}

func Update(gameState GameState) {

}
