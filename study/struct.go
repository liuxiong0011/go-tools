
package main

import (
	"fmt"
)

type Bot struct {
}

type Top struct {
	Bot
	data int
}

type Top2 struct {
	Bot
	data int
}

type Top3 struct {
	data int
	Bot
}

type Top4 struct {
	data int
	Top
	Bot
}

func (b *Bot) print() {
	fmt.Println("printBot")
}

func (b *Top) print() {
	fmt.Println("printTop")
}

func main() {
	t := new(Top)
	t.data = 1 
	t.print()

	t2 := new(Top2)
	t2.data = 2 
	t2.print()

	t3 := new(Top3)
	t3.data = 3 
	t3.print()

	// t4 := new(Top4)
	// t4.data = 4 
	// t4.print()// compile error, because both Top and bot have print()
}