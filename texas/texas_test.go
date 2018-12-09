package main

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T)  {
	card := &Card{
		//CurrentCard:"ThAs9d5c9s",
		CurrentCard:"TsAs9s5s9s",
	}
	card.SortCurrentCard()
	card.CheckCardLevel()
	fmt.Println(card, '9', 9)
}
