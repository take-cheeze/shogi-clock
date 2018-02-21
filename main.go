package main

import (
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
	termbox.Init()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	StartDisplay()
	dispL := NewDisplay(LEFT)
	dispR := NewDisplay(RIGHT)
	dispL.Print(612)
	dispR.Print(1231)
	termbox.Flush()
	time.Sleep(5 * time.Second)
	EndDisplay()
	termbox.Close()
	//fmt.Println("hello world.")

	// ch := make(chan int)

	// go test(ch, 1)
	// go test(ch, 2)
	// go test(ch, 3)

	// <-ch
	// <-ch
	// <-ch

	// for {
	// 	ret := <-ch
	// 	fmt.Println(ret)
	// 	if ret == 9 {
	// 		break
	// 	}
	// }
}
