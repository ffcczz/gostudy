package main

import (
	"strings"
)

const (
	LevelRoyalFlush         = 100 //皇家同花顺
	LevelStragStraightFlush = 99  //同花顺
	LevelFourOfKind         = 98  //四条
	LevelFullHouse          = 97  //葫芦
	LevelFlush              = 96  //同花
	LevelStraight           = 95  //顺子
	LevelThreeOfKind        = 94  //三条
	LevelTwoPairs           = 93  //两对
	LevelOnePair            = 92  //一对
	LevelOvercard           = 91  //单张最大
	ResultFirst             = 1   //
	ResultSecond            = 2
	ResultDogFall           = 0
    BaseCardLengh           = 5   //基本牌长
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
	'X':14,
}



type Card struct {    //五张   基础
	Level       int      // 此手牌所属等级   皇家同花顺 100  同花顺 99 以此类推
	CurrentCard string   //初始牌
	SortCard    string   //排序后的牌
	SortCardFace string  //排序后的牌面
	SortCardColor string //排序后的花色
	MaxCardFace rune     //牌面相同最多张
	SecondCardFace rune  //牌面相同第二多的张
}

type CardSeven struct { //七张
	Card
}

type CardFiveGhost struct {  //五张带癞子
	Card
}

type CardSevenGhost struct {  //七张带癞子
	CardSeven
}


// 将初始牌排序  并设置好 排序后的牌 牌面  和花色
func (card *Card) SortCurrentCard()  {
	currentCard := []rune(card.CurrentCard)
	//冒泡排序
	for i := 0; i < len(currentCard) -2 ; i+=2 {
		for j:=0; j < len(currentCard) -i -2; j+=2 {
			if FaceSortBaseMap[currentCard[j]] >= FaceSortBaseMap[currentCard[j+2]] {
				currentCard[j],currentCard[j+1], currentCard[j+2],currentCard[j+3] = currentCard[j+2],currentCard[j+3] , currentCard[j],currentCard[j+1]
			}
		}
	}
	//设置好 排序后的牌
	card.SortCard = string(currentCard)
	var sortCardFace []rune
	var sortCardColor []rune
	for i := 0; i < len(currentCard); i+=2 {
		sortCardFace = append(sortCardFace, currentCard[i])
		sortCardColor = append(sortCardColor, currentCard[i+1])
	}
	//设置好 排序后的牌面  和花色
	card.SortCardFace = string(sortCardFace)
	card.SortCardColor = string(sortCardColor)
}

// 判断所有牌是否为同一颜色(与五张带癞子一起用)
func (card *Card) IsSortCardColorSame() bool  {
	for i := 0; i < len(card.SortCardColor)-1; i++ {
		if card.SortCardColor[i] == 'n' || card.SortCardColor[i+1] == 'n' {
			continue
		}
		if card.SortCardColor[i] != card.SortCardColor[i+1] {
			return false
		}
	}
	return true
}

// 七张  判断是否为同一颜色  并修改排序后的牌面
func (card *CardSeven)IsSortCardColorSame() bool  {
	for _,svalue := range ColorBase {
		n := strings.Count(card.SortCardColor,string(svalue))
		if n >= BaseCardLengh {
			var sortCardFace string
			sortColor := card.SortCardColor
			cardFace := card.SortCardFace
			for i := 0; i< n; i++ {
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

// 五张  计算相同牌面最多张 和 相同牌面第二多的张
func (card *Card) SameCardMaxLen() (Max int, Second int)  {
	for _,svalue := range FaceSortBase {
		n := strings.Count(card.SortCardFace,string(svalue))
		if n > 0 {
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


// 给各手牌评等级
func (card *Card)CheckCardLevel()  {
	// 排序牌面
	sortCardFace := card.SortCardFace
	// 判断是否为顺子
	containsSortCardFace := strings.Contains(FaceSortBase, sortCardFace) || sortCardFace == FaceSortBaseNew
	// 判断是否为同花
	sameColor := card.IsSortCardColorSame()
	// 计算相同牌面最多张 和 相同牌面第二多的张
	Max, Second := card.SameCardMaxLen()
	if containsSortCardFace && strings.Contains(sortCardFace, "A") && sameColor { // 皇家同花顺
		card.Level = LevelRoyalFlush
	} else if containsSortCardFace && sameColor { // 同花顺
		card.Level = LevelStragStraightFlush
	} else if Max == 4 { // 四条
		card.Level = LevelFourOfKind
	} else if Max == 3 && Second == 2 { // 三带二
		card.Level = LevelFullHouse
	} else if sameColor { // 同花
		card.Level = LevelFlush
	} else if containsSortCardFace { // 顺子
		card.Level = LevelStraight
	} else if Max == 3 { //三条
		card.Level = LevelThreeOfKind
	} else if Max == 2 && Second == 2 { // 两对
		card.Level = LevelTwoPairs
	} else if  Max == 2 { //一对
		card.Level = LevelOnePair
	} else if Max == 1 { // 单张最大
		card.Level = LevelOvercard
	}
}


// 去除相同牌面   相同牌面只留一张
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
	//判断是否为同一颜色  并修改排序后的牌面
	sameColor := card.IsSortCardColorSame()
	// 排序牌面
	sortCardFace := card.SortCardFace
	containsSortCardFace := false
	// 去除相同牌面   相同牌面只留一张
	sortCardFace = SplitSameCard(sortCardFace)
	// 判断是否为顺子
	for i:=0; i< len(sortCardFace) - 4; i++ {
		if strings.Index(FaceSortBase, sortCardFace[i:i+5]) != -1{
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
	// 计算相同牌面最多张 和 相同牌面第二多的张
	Max, Second := card.SameCardMaxLen()
	/*if card.CurrentCard == "JdTd6d4dAdTs3d" {
		fmt.Println("JdTd6d4dAdTs3d", card,containsSortCardFace,sortCardFace)
	}
	if card.CurrentCard == "Kc9cJdTd6d4dAd" {
		fmt.Println("Kc9cJdTd6d4dAd", card,containsSortCardFace,sortCardFace)
	}*/
	if containsSortCardFace && strings.Contains(card.SortCardFace, "A") && sameColor { // 皇家同花顺
		card.Level = LevelRoyalFlush
	} else if containsSortCardFace && sameColor { // 同花顺
		card.Level = LevelStragStraightFlush
	} else if Max == 4 { // 四条
		// 将排序后的牌面只保留基本长度张数
		maxIndex := strings.Index(card.SortCardFace, string(card.MaxCardFace))
		if card.SortCardFace[0:1] == string(card.MaxCardFace) {
			card.SortCardFace = card.SortCardFace[maxIndex:maxIndex+5]
		} else{
			card.SortCardFace = card.SortCardFace[maxIndex:maxIndex+4] + card.SortCardFace[0:1]
			card.SecondCardFace =[]rune(card.SortCardFace[0:1])[0]
		}
		card.Level = LevelFourOfKind
	} else if Max == 3 && (Second == 2 || Second == 3){ // 三带二
		// 将排序后的牌面只保留基本长度张数
		maxIndex := strings.Index(card.SortCardFace, string(card.MaxCardFace))
		secondIndex := strings.Index(card.SortCardFace, string(card.SecondCardFace))
		card.SortCardFace = card.SortCardFace[maxIndex:maxIndex+3] + card.SortCardFace[secondIndex:secondIndex+2]
		card.Level = LevelFullHouse
	} else if sameColor { // 同花
		card.SortCardFace = card.SortCardFace[0:BaseCardLengh]
		card.Level = LevelFlush
	} else if containsSortCardFace { // 顺子
		card.SortCardFace = card.SortCardFace[0:BaseCardLengh]
		card.Level = LevelStraight
	} else if Max == 3 { //三条
		// 将排序后的牌面只保留基本长度张数
		maxIndex := strings.Index(card.SortCardFace, string(card.MaxCardFace))
		if maxIndex == 0 || maxIndex == 1{
			card.SortCardFace = card.SortCardFace[0:BaseCardLengh]
		} else {
			card.SortCardFace = card.SortCardFace[0:2] + card.SortCardFace[maxIndex:maxIndex+3]
		}
		card.Level = LevelThreeOfKind
	} else if Max == 2 && Second == 2 { // 两对
		// 将排序后的牌面只保留基本长度张数
		maxIndex := strings.Index(card.SortCardFace, string(card.MaxCardFace))
		secondIndex := strings.Index(card.SortCardFace, string(card.SecondCardFace))
		sortCardBase := card.SortCardFace[maxIndex:maxIndex+2] + card.SortCardFace[secondIndex:secondIndex+2]
		sortCardFace := strings.Replace(card.SortCardFace, string(card.MaxCardFace), "", -1)
		sortCardFace = strings.Replace(sortCardFace, string(card.SecondCardFace), "", -1)
		card.SortCardFace = sortCardBase + sortCardFace[0:1]
		//fmt.Println(maxIndex,secondIndex,sortCardBase,sortCardFace,card)
		card.Level = LevelTwoPairs
	} else if  Max == 2 { //一对
		// 将排序后的牌面只保留基本长度张数
		maxIndex := strings.Index(card.SortCardFace, string(card.MaxCardFace))
		//fmt.Println(card.SortCardFace, card.MaxCardFace, maxIndex)
		sortCardBase := card.SortCardFace[maxIndex:maxIndex+2]
		sortCardFace := strings.Replace(card.SortCardFace, string(card.MaxCardFace), "", -1)
		card.SortCardFace = sortCardBase + sortCardFace[0:3]
		card.Level = LevelOnePair
	} else if Max == 1 { // 单张最大
		// 将排序后的牌面只保留基本长度张数
		card.SortCardFace = card.SortCardFace[0:5]
		card.Level = LevelOvercard
	}
	/*if card.CurrentCard == "JdTd6d4dAdTs3d" {
		fmt.Println("JdTd6d4dAdTs3d", card,)
	}
	if card.CurrentCard == "Kc9cJdTd6d4dAd" {
		fmt.Println("Kc9cJdTd6d4dAd", card,)
	}*/
}

func (card *CardFiveGhost) CheckCardLevel()  {
	if strings.Index(card.SortCardFace, "X") != -1 {
		sortCardFace := card.SortCardFace
		containsSortCardFace := false
		//containsSortCardFace := strings.Contains(FaceSortBase, sortCardFace) || sortCardFace == FaceSortBaseNew  //???
		//判断是否为顺子
		for i := 0; i < len(FaceSortBase) - 4; i++ {
			count := 0
			for z := 0; z < len(sortCardFace); z++ {
				if result := strings.Count(FaceSortBase[i:i+5], sortCardFace[z:z+1]);result == 1 && strings.Count(sortCardFace, sortCardFace[z:z+1]) == 1{
					count += 1
					/*if card.CurrentCard == "2c3hXn5s5d" {
						fmt.Println("2c3hXn5s5d", card,result,FaceSortBase[i:i+5],sortCardFace[z:z+1],count)
					}*/
				}
			}
			if count == 4 {
				containsSortCardFace = true
				card.SortCardFace = FaceSortBase[i:i+5]
				break
			}
		}
		sameColor := card.IsSortCardColorSame()
		Max, Second := card.SameCardMaxLen()
		/*if card.CurrentCard == "2c3hXn5s5d" {
			fmt.Println("2c3hXn5s5d", card)
		}*/

		if containsSortCardFace && strings.Contains(sortCardFace, "A") && sameColor { // 皇家同花顺
			card.Level = LevelRoyalFlush
		} else if containsSortCardFace && sameColor { // 同花顺
			card.Level = LevelStragStraightFlush
		} else if Max == 3 { // 四条
			card.Level = LevelFourOfKind
		} else if Max == 2 && Second == 2 { // 三带二
			card.Level = LevelFullHouse
		} else if sameColor { // 同花
			card.Level = LevelFlush
		} else if containsSortCardFace { // 顺子
			card.Level = LevelStraight
		} else if Max == 2 { //三条
			card.Level = LevelThreeOfKind
		} else if  Max == 1 { //一对
			card.SortCardFace = card.SortCardFace[0:1] + card.SortCardFace[0:1] + card.SortCardFace[1:4]
			card.Level = LevelOnePair
		}
		/*if card.CurrentCard == "2c3hXn5s5d" {
			fmt.Println("2c3hXn5s5d", card)
		}*/
	} else {
		card.Card.CheckCardLevel()
	}
}

func (card *CardSevenGhost)IsSortCardColorSame() bool  {
	for _,svalue := range ColorBase {
		n := strings.Count(card.SortCardColor,string(svalue))
		if n >= BaseCardLengh -1 {
			var sortCardFace string
			sortColor := card.SortCardColor
			cardFace := card.SortCardFace
			for i := 0; i< n; i++ {
				//6s7sQsQhKhKs5s
				index := strings.Index(sortColor, string(svalue))
				sortCardFace += cardFace[index:index+1]
				sortColor = sortColor[index+1:]
				cardFace = cardFace[index+1:]
			}
			card.SortCardFace = sortCardFace
			if strings.Index(card.CurrentCard, "X") != -1 {
				card.SortCardFace += cardFace[len(cardFace)-1:]
			}
			return true
		}
	}
	return false
}


func (card *CardSevenGhost) CheckCardLevel()  {
	if strings.Index(card.CurrentCard, "X") != -1 {  //带癞子

		Max, Second := card.SameCardMaxLen()
		//判断是否可能为同花顺，如果可以为同花顺  则将排序后的牌直接置为同花顺的牌
		sameColor := card.IsSortCardColorSame()
		sortCardFace := card.SortCardFace
		containsSortCardFace := false
		//去除同牌面牌  判断是否为顺子
		sortCardFace = SplitSameCard(sortCardFace)
		/*for i:=0; i< len(sortCardFace) - 4; i++ {
			if strings.Contains(FaceSortBase, sortCardFace[i:i+4]) {
				indexFind := strings.Index(FaceSortBase, sortCardFace[i:i+4])
				if indexFind == 0 {
					card.SortCardFace = sortCardFace[i:i+4] + FaceSortBase[indexFind+BaseCardLengh-1:indexFind+BaseCardLengh]
				} else {
					card.SortCardFace = FaceSortBase[indexFind+BaseCardLengh-1:indexFind+BaseCardLengh] + sortCardFace[i:i+4]
				}
				containsSortCardFace = true
				break
			}
		}*/
		for i := 0; i < len(FaceSortBase) - 4; i++ {
			count := 0
			for z := 0; z < len(sortCardFace); z++ {
				if result := strings.Count(FaceSortBase[i:i+5], sortCardFace[z:z+1]);result == 1 && strings.Count(sortCardFace, sortCardFace[z:z+1]) == 1{
					count += 1
				}
			}
			if count == 4 {
				containsSortCardFace = true
				if Max <= 2 && Second < 2 {
					card.SortCardFace = FaceSortBase[i:i+5]
				}
				break
			}
		}
		if !containsSortCardFace {
			count := 0
			for z := 0; z < len(sortCardFace); z++ {
				if result := strings.Count(FaceSortBaseNew, sortCardFace[z:z+1]);result == 1 && strings.Count(sortCardFace, sortCardFace[z:z+1]) == 1{
					count += 1
				}
			}
			if count == 4 {
				containsSortCardFace = true
				if Max <= 2 && Second < 2 {
					card.SortCardFace = FaceSortBaseNew
				}
			}
		}
		/*if card.CurrentCard == "XnAd5s2h3d9hQs" {
			fmt.Println("XnAd5s2h3d9hQs", card,)
		}
		if card.CurrentCard == "6h6cXnAd5s2h3d" {
			fmt.Println("6h6cXnAd5s2h3d", card,)
		}*/
	/*	if !containsSortCardFace {
			for i:=0; i< len(sortCardFace) - 4; i++ {
				//if sortCardFace[i:i+5] == FaceSortBaseNew || sortCardFace[0:1]+sortCardFace[len(sortCardFace)-4:] == FaceSortBaseNew {
				if sortCardFace[i:i+5] == FaceSortBaseNew || sortCardFace[0:1]+sortCardFace[len(sortCardFace)-4:] == FaceSortBaseNew {
					card.SortCardFace = FaceSortBaseNew
					containsSortCardFace = true
					break
				}
			}
		}*/
		if containsSortCardFace && strings.Contains(card.SortCardFace, "A") && sameColor { // 皇家同花顺
			card.Level = LevelRoyalFlush
		} else if containsSortCardFace && sameColor { // 同花顺
			card.Level = LevelStragStraightFlush
		} else if Max == 3 { // 四条
			card.Level = LevelFourOfKind
		} else if Max == 2 && Second == 2 { // 三带二
			maxIndex := strings.Index(card.SortCardFace, string(card.MaxCardFace))
			secondIndex := strings.Index(card.SortCardFace, string(card.SecondCardFace))
			//fmt.Println("card.SortCardFace 三带二", card.SortCardFace,card.CurrentCard)
			card.SortCardFace = card.SortCardFace[maxIndex:maxIndex+2] + card.SortCardFace[secondIndex:secondIndex+2] + card.SortCardFace[len(card.SortCardFace)-1:]
			//fmt.Println("card.SortCardFace 三带二 2 ", card.SortCardFace,card.CurrentCard)
			card.Level = LevelFullHouse
		} else if sameColor { // 同花
			card.SortCardFace = "A" + card.SortCardFace[0:BaseCardLengh-1]
			card.Level = LevelFlush
		} else if containsSortCardFace { // 顺子
			card.Level = LevelStraight
		} else if Max == 2 { //三条
			maxIndex := strings.Index(card.SortCardFace, string(card.MaxCardFace))
			if maxIndex == 0 || maxIndex == 1{
				card.SortCardFace = card.SortCardFace[0:BaseCardLengh-1]+ card.SortCardFace[len(card.SortCardFace)-1:]
			} else {
				card.SortCardFace = card.SortCardFace[0:2] + card.SortCardFace[maxIndex:maxIndex+2] +  card.SortCardFace[len(card.SortCardFace)-1:]
			}
			card.Level = LevelThreeOfKind
	/*	} else if Max == 2 && Second == 2 { // 两对
			maxIndex := strings.Index(card.SortCardFace, string(card.MaxCardFace))
			secondIndex := strings.Index(card.SortCardFace, string(card.SecondCardFace))
			sortCardBase := card.SortCardFace[maxIndex:maxIndex+2] + card.SortCardFace[secondIndex:secondIndex+2]
			sortCardFace := strings.Replace(card.SortCardFace, string(card.MaxCardFace), "", -1)
			sortCardFace = strings.Replace(sortCardFace, string(card.SecondCardFace), "", -1)
			card.SortCardFace = sortCardBase + sortCardFace[0:1]
			//fmt.Println(maxIndex,secondIndex,sortCardBase,sortCardFace,card)
			card.LevelRoyalFlush = LevelTwoPairs*/
		} else if  Max == 1 { //一对
			maxIndex := strings.Index(card.SortCardFace, string(card.MaxCardFace))
			//fmt.Println(card.SortCardFace, card.MaxCardFace, maxIndex)
			sortCardBase := card.SortCardFace[maxIndex:maxIndex+1]
			sortCardFace := strings.Replace(card.SortCardFace, string(card.MaxCardFace), "", -1)
			card.SortCardFace = sortCardBase +sortCardBase + sortCardFace[0:3]
			card.Level = LevelOnePair
		}
		/*if card.CurrentCard == "XnAd5s2h3d9hQs" {
			fmt.Println("XnAd5s2h3d9hQs", card,)
		}
		if card.CurrentCard == "3c8h6s6c6h6dKc" {
			fmt.Println("3c8h6s6c6h6dKc", card,)
		}*/
	}else {      // 不带癞子
		card.CardSeven.CheckCardLevel()
	}
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
	case LevelStragStraightFlush, LevelStraight: // 同花顺,顺子   比较第一张
		if cardOne.SortCardFace == FaceSortBaseNew || cardTwo.SortCardFace == FaceSortBaseNew {
			if cardOne.SortCardFace == FaceSortBaseNew &&  cardTwo.SortCardFace != FaceSortBaseNew {
				return ResultSecond
			} else if cardOne.SortCardFace != FaceSortBaseNew &&  cardTwo.SortCardFace == FaceSortBaseNew {
				return ResultFirst
			} else {
				return ResultDogFall
			}
			//CompareAllCardFace([]rune(cardOne.SortCardFace),[]rune(cardTwo.SortCardFace))
		}
		if FaceSortBaseMap[cardOneSortFace[0]] < FaceSortBaseMap[cardTwoSortFace[0]] {
			return ResultFirst
		} else if FaceSortBaseMap[cardOneSortFace[0]] > FaceSortBaseMap[cardTwoSortFace[0]] {
			return ResultSecond
		} else {
			return ResultDogFall
		}
	case LevelFourOfKind, LevelFullHouse, LevelThreeOfKind, LevelTwoPairs: //三条，两对 四条、三带二  比较同牌面张和单张
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
				//fmt.Println("cardOne.SortCardFace",cardOne.SortCardFace,cardOne.CurrentCard,"cardTwo.SortCardFace",cardTwo.SortCardFace,cardTwo.CurrentCard)
				sortCardFace := cardOne.SortCardFace
				sortCardFace = strings.Replace(sortCardFace, string(cardOne.MaxCardFace), "", -1)
				sortCardFace = strings.Replace(sortCardFace, string(cardOne.SecondCardFace), "", -1)
				sortCardFace = strings.Replace(sortCardFace, "X", "", -1)
				sortCardFaceTwo := cardTwo.SortCardFace
				sortCardFaceTwo = strings.Replace(sortCardFaceTwo, string(cardTwo.MaxCardFace), "", -1)
				sortCardFaceTwo = strings.Replace(sortCardFaceTwo, string(cardTwo.SecondCardFace), "", -1)
				sortCardFaceTwo = strings.Replace(sortCardFaceTwo, "X","",-1)
				if len(sortCardFace) > len(sortCardFaceTwo) {
					sortCardFaceTwo = sortCardFaceTwo[0:len(sortCardFace)]
				} else if len(sortCardFace) < len(sortCardFaceTwo){
					sortCardFace = sortCardFace[0:len(sortCardFaceTwo)]
				}
				//fmt.Println("sortCardFace",sortCardFace,"sortCardFaceTwo",sortCardFaceTwo)
				return CompareAllCardFace([]rune(sortCardFace),[]rune(sortCardFaceTwo))
			}
		}
	case LevelOvercard, LevelFlush: // 单张最大，同花  比较所有牌面大小
		if len(cardOneSortFace) > len(cardTwoSortFace) {
			cardTwoSortFace = cardTwoSortFace[0:len(cardOneSortFace)]
		} else if len(cardOneSortFace) < len(cardTwoSortFace){
			cardOneSortFace = cardOneSortFace[0:len(cardTwoSortFace)]
		}
		return CompareAllCardFace(cardOneSortFace, cardTwoSortFace)
	case LevelOnePair: // 一对
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
	//fmt.Println(one,two,len(one),len(two))
	for i := 0; i < len(one); i++ {
		if FaceSortBaseMap[one[i]] < FaceSortBaseMap[two[i]] {
			//fmt.Println(FaceSortBaseMap[one[i]],FaceSortBaseMap[two[i]])
			return ResultFirst
		} else if FaceSortBaseMap[one[i]] > FaceSortBaseMap[two[i]]{
			return ResultSecond
		} else {
			//fmt.Println(i,"aaaa",FaceSortBaseMap[one[i]],"bbbb",FaceSortBaseMap[two[i]])
			if len(one) > 1 {
				//fmt.Println(i,"aaaabbbb",FaceSortBaseMap[one[i]],"bbbb",FaceSortBaseMap[two[i]])
				return CompareAllCardFace(one[i+1:],two[i+1:])
			}else {
				return ResultDogFall
			}
		}
	}
	return ResultDogFall
}