package main

import (
	"context"
	"time"
)

const COUNT_DOWN_SEC = 1

type Player struct {
	sec  int
	c    chan bool
	side int
}

func (player *Player) Notify(kind int) {
	switch kind {
	case STOP:
		player.c <- false
	case LOSE:
		player.c <- true
	}
}

func NewPlayer(side int, sec int, display *Display) *Player {
	var player Player
	player.sec = sec
	player.side = side
	player.c = make(chan bool)
	display.Print(player.side, player.sec)
	return &player
}

func (player *Player) Turn(display *Display, button *Button) bool {
	t := time.NewTicker(COUNT_DOWN_SEC * time.Second)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	display.BlinkOn(player.side)
	defer display.BlinkOff(player.side)

	button.Start(ctx)
	for {
		select {
		case lose := <-player.c:
			return lose
		case <-t.C:
			player.sec--
			display.Print(player.side, player.sec)
			if player.sec == 0 {
				return true
			}
		}
	}
}
