package main

import (
	"fmt"
	"math/rand"
	"time"
)

const cardFaceNumber = 0
const cardZone = 1
const (
	deck = iota
	bin
)
const cardPos = 2

var state []*[]int

func zeroTransformation(state []*[]int, cards []*[]int, transform bool) (transformed bool) {
	card := *(cards[0])
	if (*card)[cardZone] == deck && (*card)[cardPos] == maxPos(state, deck) {
		if transform {
			(*card)[cardZone] = bin
			(*card)[cardPos] = maxPos(state, bin) + 1
		}
		transformed = true
	}

	transformed = false
	return
}

func canTransformation(state []*[]int, transformation func([]*[]int, []*[]int, bool) bool) (cards [][]int) {
	for _, card := range state {
		if transformation(*card, false) {
			cards = append(cards, card)
		}
		return
	}
	return
}

func maxPos(state []*[]int, zone int) (pos int) {
	pos = -1
	for _, card := range state {
		if card[cardZone] == zone {
			if card[cardPos] > pos {
				pos = card[cardPos]
			}
		}
	}
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		randNum := rand.Intn(2)
		card := []int{
			randNum,
			0,
			maxPos(state, 0) + 1,
		}
		fmt.Println(card)
		state = append(state, *card)
	}
	t := canTransformation(state, zeroTransformation)
	fmt.Println(t)
}
