package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	termbox "github.com/nsf/termbox-go"
)

type Output interface {
	Print(sec int)
	Blink()
	Off()
}

func test(ch chan int, no int) {
	t := time.NewTicker(4 * time.Millisecond)
	for i := 0; i < 10; i++ {
		<-t.C
		fmt.Println("test" + strconv.Itoa(no))
	}
	ch <- 1
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	termbox.Init()
	defer termbox.Close()
	disp := NewDisplay(ctx)
	disp.Print(LEFT, 612)
	disp.Print(RIGHT, 1231)
	disp.BlinkOn(LEFT)
	time.Sleep(5 * time.Second)
	disp.BlinkOn(RIGHT)
	disp.BlinkOff(LEFT)
	time.Sleep(5 * time.Second)
}
