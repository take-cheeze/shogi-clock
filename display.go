package main

import (
	"context"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	LEFT = iota
	RIGHT
	SIDES
)

type Time struct {
	min   int
	sec   int
	blink bool
	disp  bool
}
type Display struct {
	times [SIDES]Time
}

const Y_PER_DISIT = 7
const X_PER_DISIT = 5
const X_PADDING = 1
const X_MARGIN = 7
const Y_MARGIN = 7
const DIV = 10
const CHARCTOR = 'a'

const BLINK_PERIOD_MS = 250

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

func NewDisplay(ctx context.Context) *Display {
	var display Display
	for side := 0; side < SIDES; side++ {
		display.times[side].min = 0
		display.times[side].sec = 0
		display.times[side].blink = false
		display.times[side].disp = false
	}
	go func(display *Display, ctx context.Context) {

		t := time.NewTicker(BLINK_PERIOD_MS * time.Millisecond)

		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
				offset := 0
				printBothTime(&display.times, offset)
				termbox.Flush()
			case display := <-dispRequest:
				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
				offset := 0
				printBothTime(&display.times, offset)
				termbox.Flush()
			}
		}
	}(&display, ctx)
	return &display
}

func (display *Display) Print(side int, sec int) {
	display.times[side].min = sec / 60
	display.times[side].sec = sec % 60
	display.times[side].disp = true
	dispRequest <- display
}

func (display *Display) BlinkOn(side int) {
	display.times[side].blink = true
}

func (display *Display) BlinkOff(side int) {
	display.times[side].blink = false
}

func printBothTime(times *[SIDES]Time, offset int) int {
	for side := 0; side < SIDES; side++ {
		offset = printTime(&times[side], offset)
		offset++
	}
	return offset
}
func printTime(time *Time, offset int) int {
	if time.blink == true {
		time.disp = !(time.disp)
	} else {
		time.disp = true
	}
	if time.disp == true {
		offset = printTowDisit(time.min, offset)
		offset = printColon(offset)
		offset = printTowDisit(time.sec, offset)
	} else {
		offset += 2 + 1 + 2
	}
	return offset
}

func printTowDisit(num int, offset int) int {
	for disit := 0; disit < 2; disit++ {
		output := num / DIV % DIV
		offset = printDisit(output, offset)
		num *= DIV
	}
	return offset
}

func printDisit(num int, offset int) int {
	offsetX := offset*(X_PADDING+X_PER_DISIT) + X_MARGIN
	offsetY := Y_MARGIN
	for x := 0; x < X_PER_DISIT; x++ {
		for y := 0; y < Y_PER_DISIT; y++ {
			if NUM_MAP[num][y][x] {
				termbox.SetCell(x+offsetX, y+offsetY, CHARCTOR, termbox.ColorDefault, termbox.ColorDefault)
			}
		}
	}
	return offset + 1
}

func printColon(offset int) int {
	offsetX := offset*(X_PADDING+X_PER_DISIT) + X_MARGIN
	offsetY := Y_MARGIN
	for x := 0; x < X_PER_DISIT; x++ {
		for y := 0; y < Y_PER_DISIT; y++ {
			if COLON_MAP[y][x] {
				termbox.SetCell(x+offsetX, y+offsetY, CHARCTOR, termbox.ColorDefault, termbox.ColorDefault)
			}
		}
	}
	return offset + 1
}
