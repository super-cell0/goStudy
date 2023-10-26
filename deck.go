package main

import "fmt"

type deck []string

// 打印循环
func (index deck) printDeck() {
	for i, card := range index {
		fmt.Println(i, card)
	}

}
