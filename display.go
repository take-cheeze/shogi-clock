package main

import (
	"github.com/nsf/termbox-go"
)

type Display struct {
	side   int
	offset int
	min    int
	sec    int
}

const (
	LEFT = iota
	RIGHT
)

const Y_PER_DISIT = 7
const X_PER_DISIT = 5
const X_PADDING = 1
const OFFSET_DISITS = 6
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
	display.offset = side * (X_PER_DISIT + X_PADDING) * OFFSET_DISITS
	display.min = 0
	display.sec = 0

	return &display
}

func (display *Display) Print(sec int) {
	display.min = sec / 60
	display.sec = sec % 60
	dispRequest <- display
}

func (display *Display) Blink() {

}

func (display *Display) Off() {

}
