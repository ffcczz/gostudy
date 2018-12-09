package main

import (
	"fmt"
	"strings"
)

const (
	LevelFirst = 100
	LevelSecond = 99
	LevelThird = 98
	LevelFourth = 97
	LevelFifth = 96
	LevelSixth = 95
	LevelSeventh = 94
	LevelEighth = 93
	LevelNinth = 92
	LevelTenth = 91
)

var FaceBase = "23456789TJQKA"
var FaceSortBase = "AKQJT98765432"
var FaceSortBaseMap = map[rune]rune{
	'A':1,
	'K':2,
	'Q':3,
	'J':4,
	'T':5,
	'9':6,
	'8':7,
	'7':8,
	'6':9,
	'5':10,
	'4':11,
	'3':12,
	'2':13,
}
type Card struct {
	Level       int   // 此手牌所属等级   皇家同花顺 100  同花顺 99 以此类推
	CurrentCard string
	SortCard    string
	SortCardFace string
	SortCardColor string
	MaxCardFace rune
	SecondCardFace rune
}

func (card *Card) SortCurrentCard()  {
	currentCard := []rune(card.CurrentCard)
	for i := 0; i < len(currentCard) -2 ; i+=2 {
		for j:=0; j < len(currentCard) -i -2; j+=2 {
			//ThAs9d5c9s
			if FaceSortBaseMap[currentCard[j]] >= FaceSortBaseMap[currentCard[j+2]] {
				fmt.Println(FaceSortBaseMap[currentCard[j]] , currentCard[j])
				fmt.Println(FaceSortBaseMap[currentCard[j+2]], currentCard[j+2] )
				currentCard[j],currentCard[j+1], currentCard[j+2],currentCard[j+3] = currentCard[j+2],currentCard[j+3] , currentCard[j],currentCard[j+1]
			}
		}
	}
	card.SortCard = string(currentCard)
	var sortCardFace []rune
	var sortCardColor []rune
	for i := 0; i < len(currentCard); i+=2 {
		sortCardFace = append(sortCardFace, currentCard[i])
		sortCardColor = append(sortCardColor, currentCard[i+1])
	}
	card.SortCardFace = string(sortCardFace)
	card.SortCardColor = string(sortCardColor)
}


func (card *Card) IsSortCardColorSame() bool  {
	for i := 0; i < len(card.SortCardColor)-1; i++ {
		if card.SortCardColor[i] != card.SortCardColor[i+1] {
			return false
		}
	}
	return true
}

func (card *Card) SameCardMaxLen() (Max int, Second int)  {
	for _,svalue := range FaceSortBase {
		n := strings.Count(card.SortCardFace,string(svalue))
		if Max <= n {
			Max,Second = n,Max
			card.MaxCardFace, card.SecondCardFace = svalue, card.MaxCardFace
		}
	}
	return
}


func (card *Card)CheckCardLevel()  {
	sortCardFace := card.SortCardFace
	containsSortCardFace := strings.Contains(FaceSortBase, sortCardFace)
	sameColor := card.IsSortCardColorSame()
	Max, Second := card.SameCardMaxLen()
	if containsSortCardFace && strings.Contains(sortCardFace, "A") && sameColor { // 皇家同花顺
		card.Level = LevelFirst
	} else if containsSortCardFace && sameColor { // 同花顺
		card.Level = LevelSecond
	} else if Max == 4 { // 四条
		card.Level = LevelThird
	} else if Max == 3 && Second == 2 { // 三带二
		card.Level = LevelFourth
	} else if sameColor { // 同花
		card.Level = LevelFifth
	} else if containsSortCardFace { // 顺子
		card.Level = LevelSixth
	} else if Max == 3 { //三条
		card.Level = LevelSeventh
	} else if Max == 2 && Second == 2 { // 两对
		card.Level = LevelEighth
	} else if  Max == 2 { //一对
		card.Level = LevelNinth
	} else if Max == 1 { // 单张最大
		card.Level = LevelTenth
	}

}




func CompareCard(cardOne *Card, cardTwo *Card) (win int) {
	if cardOne.Level > cardTwo.Level {
		return 1
	} else if cardOne.Level < cardTwo.Level {
		return 2
	}

	cardOneSortFace := []rune(cardOne.SortCardFace)
	cardTwoSortFace := []rune(cardTwo.SortCardFace)
	switch cardOne.Level {
	case LevelSecond, LevelSixth:  // 同花顺,顺子   比较第一张
		if FaceSortBaseMap[cardOneSortFace[0]] < FaceSortBaseMap[cardTwoSortFace[0]] {
			return 1
		} else if FaceSortBaseMap[cardOneSortFace[0]] > FaceSortBaseMap[cardTwoSortFace[0]] {
			return 2
		} else {
			return 3
		}
	case LevelThird, LevelFourth,LevelSeventh,LevelEighth:  //三条，两对 四条、三带二  比较同牌面张和单张
		if FaceSortBaseMap[cardOne.MaxCardFace] < FaceSortBaseMap[cardTwo.MaxCardFace] {
			return 1
		} else if FaceSortBaseMap[cardOne.MaxCardFace] > FaceSortBaseMap[cardTwo.MaxCardFace] {
			return 2
		} else {
			if FaceSortBaseMap[cardOne.SecondCardFace] < FaceSortBaseMap[cardTwo.SecondCardFace] {
				return 1
			} else if FaceSortBaseMap[cardOne.SecondCardFace] > FaceSortBaseMap[cardTwo.SecondCardFace] {
				return 2
			} else {
				return 3
			}
		}
	case LevelTenth,LevelFifth:  // 单张最大，同花  比较所有牌面大小
		return CompareAllCardFace(cardOneSortFace, cardTwoSortFace)
	case LevelNinth:  // 一对
		if FaceSortBaseMap[cardOne.MaxCardFace] < FaceSortBaseMap[cardTwo.MaxCardFace] {
			return 1
		} else if FaceSortBaseMap[cardOne.MaxCardFace] > FaceSortBaseMap[cardTwo.MaxCardFace] {
			return 2
		} else {
			return CompareAllCardFace(cardOneSortFace,cardTwoSortFace)
		}
	}
	return 3
}

func CompareAllCardFace(one,two []rune) int {
	for i := 0; i < len(one); i++ {
		if FaceSortBaseMap[one[i]] < FaceSortBaseMap[two[i]] {
			return 1
		} else if FaceSortBaseMap[one[i]] > FaceSortBaseMap[two[i]]{
			return 2
		} else {
			return CompareAllCardFace(one[i+1:],two[i+1:])
		}
	}
	return 3
}