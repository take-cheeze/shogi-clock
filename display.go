package main

import (
	"github.com/nsf/termbox-go"
)

type Display struct {
	side   int
	offset int
<<<<<<< HEAD
	min    int
	sec    int
=======
>>>>>>> 01836219940d7c4e9299b68abd899a083bb9b1f6
}

const (
	LEFT = iota
	RIGHT
)

const Y_PER_DISIT = 7
const X_PER_DISIT = 5
const X_PADDING = 1
<<<<<<< HEAD
const OFFSET_DISITS = 6
=======
const DISIT_PER_SIDE = 4
>>>>>>> 01836219940d7c4e9299b68abd899a083bb9b1f6
const DIV = 10

var COLON_MAP = [][]bool{
	{false, false, false, false, false},
	{false, false, true, false, false},
	{false, false, false, false, false},
	{false, false, false, false, false},
	{false, false, false, false, false},
	{false, false, true, false, false},
	{false, false, false, false, false}}

var NUM_MAP = [][][]bool{
	{{true, true, true, true, true},
		{true, false, false, false, true},
		{true, false, false, false, true},
		{true, false, false, false, true},
		{true, false, false, false, true},
		{true, false, false, false, true},
		{true, true, true, true, true}},
	{{false, true, true, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, true, true, true, false}},
	{{true, true, true, true, true},
		{false, false, false, false, true},
		{false, false, false, false, true},
		{true, true, true, true, true},
		{true, false, false, false, false},
		{true, false, false, false, false},
		{true, true, true, true, true}},
	{{true, true, true, true, true},
		{false, false, false, false, true},
		{false, false, false, false, true},
		{true, true, true, true, true},
		{false, false, false, false, true},
		{false, false, false, false, true},
		{true, true, true, true, true}},
	{{true, false, false, true, false},
		{true, false, false, true, false},
		{true, false, false, true, false},
		{true, true, true, true, true},
		{false, false, false, true, false},
		{false, false, false, true, false},
		{false, false, false, true, false}},
	{{true, true, true, true, true},
		{true, false, false, false, false},
		{true, false, false, false, false},
		{true, true, true, true, true},
		{false, false, false, false, true},
		{false, false, false, false, true},
		{true, true, true, true, true}},
	{{true, true, true, true, true},
		{true, false, false, false, false},
		{true, false, false, false, false},
		{true, true, true, true, true},
		{true, false, false, false, true},
		{true, false, false, false, true},
		{true, true, true, true, true}},
	{{true, true, true, true, true},
		{false, false, false, false, true},
		{false, false, false, false, true},
		{false, false, false, false, true},
		{false, false, false, false, true},
		{false, false, false, false, true},
		{false, false, false, false, true}},
	{{true, true, true, true, true},
		{true, false, false, false, true},
		{true, false, false, false, true},
		{true, true, true, true, true},
		{true, false, false, false, true},
		{true, false, false, false, true},
		{true, true, true, true, true}},
	{{true, true, true, true, true},
		{true, false, false, false, true},
		{true, false, false, false, true},
		{true, true, true, true, true},
		{false, false, false, false, true},
		{false, false, false, false, true},
		{true, true, true, true, true}}}

var dispRequest = make(chan *Display)
var exitRequest = make(chan bool)

func StartDisplay() {
	go func() {
	loop:
		for {
			select {
			case <-exitRequest:
				break loop
			case display := <-dispRequest:
				output := 0
				offset := display.offset
				min := display.min
				sec := display.sec

				for disit := 0; disit < 2; disit++ {
					output = min / DIV % DIV
					for x := 0; x < X_PER_DISIT; x++ {
						for y := 0; y < Y_PER_DISIT; y++ {
							if NUM_MAP[output][y][x] {
								termbox.SetCell(x+offset, y, 'a', termbox.ColorDefault, termbox.ColorDefault)
							}
						}
					}
					min *= DIV
					offset += (X_PADDING + X_PER_DISIT)
				}

				for x := 0; x < X_PER_DISIT; x++ {
					for y := 0; y < Y_PER_DISIT; y++ {
						if COLON_MAP[y][x] {
							termbox.SetCell(x+offset, y, 'a', termbox.ColorDefault, termbox.ColorDefault)
						}
					}
				}
				offset += (X_PADDING + X_PER_DISIT)

				for disit := 0; disit < 2; disit++ {
					output = sec / DIV % DIV
					for x := 0; x < X_PER_DISIT; x++ {
						for y := 0; y < Y_PER_DISIT; y++ {
							if NUM_MAP[output][y][x] {
								termbox.SetCell(x+offset, y, 'a', termbox.ColorDefault, termbox.ColorDefault)
							}
						}
					}
					sec *= DIV
					offset += (X_PADDING + X_PER_DISIT)
				}
			}
		}
	}()
}

func EndDisplay() {
	exitRequest <- true
}

func NewDisplay(side int) *Display {
	var display Display

	display.side = side
<<<<<<< HEAD
	display.offset = side * (X_PER_DISIT + X_PADDING) * OFFSET_DISITS
	display.min = 0
	display.sec = 0
=======
	display.offset = side * X_PER_DISIT * DISIT_PER_SIDE
>>>>>>> 01836219940d7c4e9299b68abd899a083bb9b1f6

	return &display
}

func (display *Display) Print(sec int) {
<<<<<<< HEAD
	display.min = sec / 60
	display.sec = sec % 60
	dispRequest <- display
=======
	offset := 0
	output := 0

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for x := 0; x < X_PER_DISIT; x++ {
		for y := 0; y < Y_PER_DISIT; y++ {
			if COLON_MAP[y][x] {
				termbox.SetCell(x+offset, y, 'a', termbox.ColorDefault, termbox.ColorDefault)
			}
		}
	}
	offset += (X_PADDING + X_PER_DISIT)

	for disit := 0; disit < 2; disit++ {
		output = sec / DIV % DIV
		for x := 0; x < X_PER_DISIT; x++ {
			for y := 0; y < Y_PER_DISIT; y++ {
				if NUM_MAP[output][y][x] {
					termbox.SetCell(x+offset, y, 'a', termbox.ColorDefault, termbox.ColorDefault)
				}
			}
		}
		sec *= DIV
		offset += (X_PADDING + X_PER_DISIT)
	}
	termbox.Flush()
>>>>>>> 01836219940d7c4e9299b68abd899a083bb9b1f6
}

func (display *Display) Blink() {

}

func (display *Display) Off() {

}
