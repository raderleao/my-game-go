package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/raderleao/mygamego/game"
)

func main() {
	g := game.NewGame()

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}

}
