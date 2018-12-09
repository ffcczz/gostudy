package main

import (
	"fmt"
	"testing"
	"time"
)

func TestSort(t *testing.T)  {
	card := &Card{
		//CurrentCard:"ThAs9d5c9s",
		//CurrentCard:"TsAs9s5s9s",
		//CurrentCard:"AsAhAdAc5s",
		//CurrentCard:"KhKsQsQh5s",
		CurrentCard:"QsQhKhKs5s",
	}
	card.SortCurrentCard()
	card.CheckCardLevel()
	fmt.Println(card, '9', 9,"K:",'K',"Q:",'Q')
}

func TestCardSeven(t *testing.T)  {
	card := &CardSeven{
		//CurrentCard:"ThAs9d5c9s",
		//CurrentCard:"TsAs9s5s9s",
		//CurrentCard:"AsAhAdAc5s",
		//CurrentCard:"KhKsQsQh5s",
		Card:Card{
			CurrentCard:"6s7sQsQhKhKs5s",
		},
	}
	card.SortCurrentCard()
	card.CheckCardLevel()
	fmt.Println(card)
}

func TestCard_SameCardMaxLen(t *testing.T) {

}

func TestCompare(t *testing.T) {
	//currentPath,_ := filepath.Abs(filepath.Dir(os.Args[0]))
	//fmt.Println(currentPath)
	var matches FiveCards
	var results FiveCards
	ReadFile("match.json", &matches)
	start := time.Now().UTC().Nanosecond()
	for i := 0; i< len(matches.Matches); i++ {
		cardBob := Card{
			CurrentCard:matches.Matches[i].Bob,
		}
		cardAlice := Card{
			CurrentCard:matches.Matches[i].Alice,
		}
		cardBob.SortCurrentCard()
		cardBob.CheckCardLevel()
		cardAlice.SortCurrentCard()
		cardAlice.CheckCardLevel()
		matches.Matches[i].Result = CompareCard(&cardAlice,&cardBob)
	}
	end := time.Now().UTC().Nanosecond()
	fmt.Println(end-start)
	ReadFile("result.json", &results)
	for i:=0; i < len(matches.Matches); i++ {
		if matches.Matches[i].Result != results.Matches[i].Result {
			fmt.Println(i,"result error", matches.Matches[i],results.Matches[i])
		}
	}
}

func TestCompareSevenCard(t *testing.T) {
	//currentPath,_ := filepath.Abs(filepath.Dir(os.Args[0]))
	//fmt.Println(currentPath)
	var matches FiveCards
	var results FiveCards
	ReadFile("seven_cards.json", &matches)
	start := time.Now().UTC().Nanosecond()
	for i := 0; i< len(matches.Matches); i++ {
		cardBob := CardSeven{
			Card:Card{
				CurrentCard:matches.Matches[i].Bob,
			},


		}
		cardAlice := CardSeven{
			Card:Card{
			CurrentCard:matches.Matches[i].Alice,
		    },
		}
		cardBob.SortCurrentCard()
		cardBob.CheckCardLevel()
		//fmt.Println("cardBob",cardBob)
		cardAlice.SortCurrentCard()
		cardAlice.CheckCardLevel()
		//fmt.Println("cardAlice",cardAlice)
		matches.Matches[i].Result = CompareCard(&cardAlice.Card,&cardBob.Card)
	}
	end := time.Now().UTC().Nanosecond()
	fmt.Println(end-start)
	ReadFile("seven_cards.result.json", &results)
	count := 0
	for i:=0; i < len(matches.Matches); i++ {
		if matches.Matches[i].Result != results.Matches[i].Result {
			count += 1
			fmt.Println(i,"result error", matches.Matches[i],results.Matches[i], count)
		}
	}
	//fmt.Println(matches)
}