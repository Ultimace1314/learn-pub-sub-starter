package main

import (
	"fmt"

	"github.com/Ultimace1314/learn-pub-sub-starter/internal/gamelogic"
	"github.com/Ultimace1314/learn-pub-sub-starter/internal/routing"
)

func handlerPause(gs *gamelogic.GameState) func(routing.PlayingState) {
	return func(ps routing.PlayingState) {
		defer fmt.Print("> ")
		gs.HandlePause(ps)
	}
}
