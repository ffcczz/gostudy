package main

import (
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
	ResultFirst = 1
	ResultSecond = 2
	ResultDogFall = 0
)

var FaceBase = "23456789TJQKA"
var FaceSortBase = "AKQJT98765432"
var FaceSortBaseNew = "A5432"
var ColorBase = "shdc"
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

type CardSeven struct {
	Card
}

type CardFiveGhost struct {
	Card
}

func (card *Card) SortCurrentCard()  {
	currentCard := []rune(card.CurrentCard)
	for i := 0; i < len(currentCard) -2 ; i+=2 {
		for j:=0; j < len(currentCard) -i -2; j+=2 {
			//ThAs9d5c9s
			if FaceSortBaseMap[currentCard[j]] >= FaceSortBaseMap[currentCard[j+2]] {
				//fmt.Println(FaceSortBaseMap[currentCard[j]] , currentCard[j])
				//fmt.Println(FaceSortBaseMap[currentCard[j+2]], currentCard[j+2] )
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

const BaseCardLengh  = 5

func (card *CardSeven)IsSortCardColorSame() bool  {
	for _,svalue := range ColorBase {
		n := strings.Count(card.SortCardColor,string(svalue))
		if n >= BaseCardLengh {
			var sortCardFace string
			sortColor := card.SortCardColor
			cardFace := card.SortCardFace
			for i := 0; i< BaseCardLengh; i++ {
				//6s7sQsQhKhKs5s
				index := strings.Index(sortColor, string(svalue))
				sortCardFace += cardFace[index:index+1]
				sortColor = sortColor[index+1:]
				cardFace = cardFace[index+1:]
			}
			card.SortCardFace = sortCardFace
			return true
		}
	}
	return false
}

func (card *Card) SameCardMaxLen() (Max int, Second int)  {
	for _,svalue := range FaceSortBase {
		n := strings.Count(card.SortCardFace,string(svalue))
		if n > 0 {
			//fmt.Println(n, svalue)
			if Max < n {
				Max,Second = n,Max
				card.MaxCardFace, card.SecondCardFace = svalue, card.MaxCardFace
			} else if Second < n {
				Second = n
				card.SecondCardFace = svalue
			}
		}
	}
	return
}


func (card *Card)CheckCardLevel()  {
	sortCardFace := card.SortCardFace
	containsSortCardFace := strings.Contains(FaceSortBase, sortCardFace) || sortCardFace == FaceSortBaseNew
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

func SplitSameCard(sortCardFace string) string  {
	for i:=0; i < len(sortCardFace) -1; i++ {
		if sortCardFace[i] == sortCardFace[i+1] {
			sortCardFace = sortCardFace[0:i] + sortCardFace[i+1:]
			sortCardFace = SplitSameCard(sortCardFace)
		}
	}
	return sortCardFace
}

func (card *CardSeven)CheckCardLevel()  {
	sameColor := card.IsSortCardColorSame()
	sortCardFace := card.SortCardFace
	containsSortCardFace := false
	sortCardFace = SplitSameCard(sortCardFace)
	for i:=0; i< len(sortCardFace) - 4; i++ {
		if strings.Contains(FaceSortBase, sortCardFace[i:i+5]) {
			card.SortCardFace = sortCardFace[i:i+5]
			containsSortCardFace = true
			break
		}
	}
	if !containsSortCardFace {
		for i:=0; i< len(sortCardFace) - 4; i++ {
			if sortCardFace[i:i+5] == FaceSortBaseNew || sortCardFace[0:1]+sortCardFace[len(sortCardFace)-4:] == FaceSortBaseNew {
				card.SortCardFace = FaceSortBaseNew
				containsSortCardFace = true
				break
			}
		}
	}
	Max, Second := card.SameCardMaxLen()
	if containsSortCardFace && strings.Contains(card.SortCardFace, "A") && sameColor { // 皇家同花顺
		card.Level = LevelFirst
	} else if containsSortCardFace && sameColor { // 同花顺
		card.Level = LevelSecond
	} else if Max == 4 { // 四条
		card.Level = LevelThird
	} else if Max == 3 && (Second == 2 || Second == 3){ // 三带二
		maxIndex := strings.Index(card.SortCardFace, string(card.MaxCardFace))
		secondIndex := strings.Index(card.SortCardFace, string(card.SecondCardFace))
		card.SortCardFace = card.SortCardFace[maxIndex:maxIndex+3] + card.SortCardFace[secondIndex:secondIndex+2]
		card.Level = LevelFourth
	} else if sameColor { // 同花
		card.Level = LevelFifth
	} else if containsSortCardFace { // 顺子
		card.Level = LevelSixth
	} else if Max == 3 { //三条
		maxIndex := strings.Index(card.SortCardFace, string(card.MaxCardFace))
		if maxIndex == 0 || maxIndex == 1{
			card.SortCardFace = card.SortCardFace[0:BaseCardLengh]
		} else {
			card.SortCardFace = card.SortCardFace[0:2] + card.SortCardFace[maxIndex:maxIndex+3]
		}
		card.Level = LevelSeventh
	} else if Max == 2 && Second == 2 { // 两对
		maxIndex := strings.Index(card.SortCardFace, string(card.MaxCardFace))
		secondIndex := strings.Index(card.SortCardFace, string(card.SecondCardFace))
		sortCardBase := card.SortCardFace[maxIndex:maxIndex+2] + card.SortCardFace[secondIndex:secondIndex+2]
		sortCardFace := strings.Replace(card.SortCardFace, string(card.MaxCardFace), "", -1)
		sortCardFace = strings.Replace(sortCardFace, string(card.SecondCardFace), "", -1)
		card.SortCardFace = sortCardBase + sortCardFace[0:1]
		//fmt.Println(maxIndex,secondIndex,sortCardBase,sortCardFace,card)
		card.Level = LevelEighth
	} else if  Max == 2 { //一对
		maxIndex := strings.Index(card.SortCardFace, string(card.MaxCardFace))
		//fmt.Println(card.SortCardFace, card.MaxCardFace, maxIndex)
		sortCardBase := card.SortCardFace[maxIndex:maxIndex+2]
		sortCardFace := strings.Replace(card.SortCardFace, string(card.MaxCardFace), "", -1)
		card.SortCardFace = sortCardBase + sortCardFace[0:3]
		card.Level = LevelNinth
	} else if Max == 1 { // 单张最大
		card.SortCardFace = card.SortCardFace[0:5]
		card.Level = LevelTenth
	}
	/*if card.CurrentCard == "5d3s2c6h4cKh8h" {
		fmt.Println("5d3s2c6h4cKh8h", card)
	}
	if card.CurrentCard == "Ah7h5d3s2c6h4c" {
		fmt.Println("Ah7h5d3s2c6h4c", card)
	}*/
}



func CompareCard(cardOne *Card, cardTwo *Card) (win int) {
	if cardOne.Level > cardTwo.Level {
		return ResultFirst
	} else if cardOne.Level < cardTwo.Level {
		return ResultSecond
	}

	cardOneSortFace := []rune(cardOne.SortCardFace)
	cardTwoSortFace := []rune(cardTwo.SortCardFace)
	switch cardOne.Level {
	case LevelSecond, LevelSixth:  // 同花顺,顺子   比较第一张
		if cardOne.SortCardFace == FaceSortBaseNew || cardTwo.SortCardFace == FaceSortBaseNew {
			if FaceSortBaseMap[cardOneSortFace[0]] < FaceSortBaseMap[cardTwoSortFace[0]] {
				return ResultSecond
			} else if FaceSortBaseMap[cardOneSortFace[0]] > FaceSortBaseMap[cardTwoSortFace[0]] {
				return ResultFirst
			} else {
				return ResultDogFall
			}
		}
		if FaceSortBaseMap[cardOneSortFace[0]] < FaceSortBaseMap[cardTwoSortFace[0]] {
			return ResultFirst
		} else if FaceSortBaseMap[cardOneSortFace[0]] > FaceSortBaseMap[cardTwoSortFace[0]] {
			return ResultSecond
		} else {
			return ResultDogFall
		}
	case LevelThird, LevelFourth,LevelSeventh,LevelEighth:  //三条，两对 四条、三带二  比较同牌面张和单张
		if FaceSortBaseMap[cardOne.MaxCardFace] < FaceSortBaseMap[cardTwo.MaxCardFace] {
			return ResultFirst
		} else if FaceSortBaseMap[cardOne.MaxCardFace] > FaceSortBaseMap[cardTwo.MaxCardFace] {
			return ResultSecond
		} else {
			if FaceSortBaseMap[cardOne.SecondCardFace] < FaceSortBaseMap[cardTwo.SecondCardFace] {
				return ResultFirst
			} else if FaceSortBaseMap[cardOne.SecondCardFace] > FaceSortBaseMap[cardTwo.SecondCardFace] {
				return ResultSecond
			} else {
				sortCardFace := cardOne.SortCardFace
				sortCardFace = strings.Replace(sortCardFace, string(cardOne.MaxCardFace), "", -1)
				sortCardFace = strings.Replace(sortCardFace, string(cardOne.SecondCardFace), "", -1)
				sortCardFaceTwo := cardTwo.SortCardFace
				sortCardFaceTwo = strings.Replace(sortCardFaceTwo, string(cardTwo.MaxCardFace), "", -1)
				sortCardFaceTwo = strings.Replace(sortCardFaceTwo, string(cardTwo.SecondCardFace), "", -1)
				return CompareAllCardFace([]rune(sortCardFace),[]rune(sortCardFaceTwo))
			}
		}
	case LevelTenth,LevelFifth:  // 单张最大，同花  比较所有牌面大小
		return CompareAllCardFace(cardOneSortFace, cardTwoSortFace)
	case LevelNinth:  // 一对
		if FaceSortBaseMap[cardOne.MaxCardFace] < FaceSortBaseMap[cardTwo.MaxCardFace] {
			return ResultFirst
		} else if FaceSortBaseMap[cardOne.MaxCardFace] > FaceSortBaseMap[cardTwo.MaxCardFace] {
			return ResultSecond
		} else {
			return CompareAllCardFace(cardOneSortFace,cardTwoSortFace)
		}
	}
	return ResultDogFall
}

func CompareAllCardFace(one,two []rune) int {
	for i := 0; i < len(one); i++ {
		if FaceSortBaseMap[one[i]] < FaceSortBaseMap[two[i]] {
			return ResultFirst
		} else if FaceSortBaseMap[one[i]] > FaceSortBaseMap[two[i]]{
			return ResultSecond
		} else {
			return CompareAllCardFace(one[i+1:],two[i+1:])
		}
	}
	return ResultDogFall
}