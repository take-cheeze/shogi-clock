package main

import (
	"context"

	termbox "github.com/nsf/termbox-go"
)

type Game struct {
	c            chan int
	firstPlayer  *Player
	secondPlayer *Player
	firstButton  *Button
	secondButton *Button
	display      *Display
}

func NewGame(firstPlayer *Player, secondPlayer *Player, firstButton *Button, secondButton *Button, display *Display) *Game {
	var game Game
	game.firstPlayer = firstPlayer
	game.secondPlayer = secondPlayer
	game.firstButton = firstButton
	game.secondButton = secondButton
	game.display = display
	game.c = make(chan int)
	return &game
}

func (game *Game) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for {
		NewButton(
			[]Reserve{
				Reserve{termbox.KeyCtrlP, PLAY, game},
				Reserve{termbox.KeyCtrlQ, QUIT, game},
			},
		).Start(ctx)
		if <-game.c == QUIT {
			break
		}
		for {
			if game.firstPlayer.Turn(game.display, game.firstButton) {
				break
			}
			if game.secondPlayer.Turn(game.display, game.secondButton) {
				break
			}
		}
	}
}

func (game *Game) Notify(kind int) {
	game.c <- kind
}
