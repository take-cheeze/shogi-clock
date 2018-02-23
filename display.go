package main

import (
	"context"

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
const CHARCTOR = 'a'

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

func StartDisplay(ctx context.Context) {
	go func(ctx context.Context) {
		termbox.Init()
		defer termbox.Close()

		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

		for {
			select {
			case <-ctx.Done():
				return
			case display := <-dispRequest:
				offset := display.offset
				min := display.min
				sec := display.sec

				for disit := 0; disit < 2; disit++ {
					output := min / DIV % DIV
					display.printDisit(output, offset)
					min *= DIV
					offset++
				}

				display.printColon(offset)

				offset++

				for disit := 0; disit < 2; disit++ {
					output := sec / DIV % DIV
					display.printDisit(output, offset)
					sec *= DIV
					offset++
				}
				termbox.Flush()
			}
		}
	}(ctx)
}

func NewDisplay(side int) *Display {
	var display Display

	display.side = side
	display.offset = side * OFFSET_DISITS
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

func (display *Display) printDisit(num int, offset int) {
	offset *= (X_PADDING + X_PER_DISIT)
	for x := 0; x < X_PER_DISIT; x++ {
		for y := 0; y < Y_PER_DISIT; y++ {
			if NUM_MAP[num][y][x] {
				termbox.SetCell(x+offset, y, CHARCTOR, termbox.ColorDefault, termbox.ColorDefault)
			}
		}
	}
}

func (display *Display) printColon(offset int) {
	offset *= (X_PADDING + X_PER_DISIT)
	for x := 0; x < X_PER_DISIT; x++ {
		for y := 0; y < Y_PER_DISIT; y++ {
			if COLON_MAP[y][x] {
				termbox.SetCell(x+offset, y, CHARCTOR, termbox.ColorDefault, termbox.ColorDefault)
			}
		}
	}
}
