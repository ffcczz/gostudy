package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

var cores = runtime.NumCPU()

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

func TestCompareFiveCard(t *testing.T) {
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
	CompareCardWithFile("match.json","result.json", CardSeven{},CardSeven{})

}

func TestCompareSevenCard(t *testing.T) {
	/*var matches FiveCards
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
		cardAlice.SortCurrentCard()
		cardAlice.CheckCardLevel()
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
	}*/

	CompareCardWithFile("seven_cards.json","seven_cards.result.json", CardSeven{},CardSeven{})
}

func CompareCardWithFile(inputFilePath, resultFilePath string,cardBob,cardAlice CardSeven)  {
	var matches FiveCards
	var results FiveCards
	ReadFile(inputFilePath, &matches)
	start := time.Now().UTC().Nanosecond()
	for i := 0; i< len(matches.Matches); i++ {
		cardBob.CurrentCard = matches.Matches[i].Bob
		cardAlice.CurrentCard = matches.Matches[i].Alice
		cardBob.SortCurrentCard()
		cardBob.CheckCardLevel()
		cardAlice.SortCurrentCard()
		cardAlice.CheckCardLevel()
		matches.Matches[i].Result = CompareCard(&cardAlice.Card,&cardBob.Card)
	}
	end := time.Now().UTC().Nanosecond()
	fmt.Println(end-start)
	ReadFile(resultFilePath, &results)
	count := 0
	for i:=0; i < len(matches.Matches); i++ {
		if matches.Matches[i].Result != results.Matches[i].Result {
			count += 1
			fmt.Println(i,"result error", matches.Matches[i],results.Matches[i], count)
		}
	}
}

func TestCompareFiveCardWithGhost(t *testing.T)  {
	var matches FiveCards
	var results FiveCards
	ReadFile("five_cards_with_ghost.json", &matches)
	start := time.Now().UTC().Nanosecond()
	for i := 0; i< len(matches.Matches); i++ {
		cardBob := CardFiveGhost{
			Card:Card{
				CurrentCard:matches.Matches[i].Bob,
			},
		}
		cardAlice := CardFiveGhost{
			Card:Card{
				CurrentCard:matches.Matches[i].Alice,
			},
		}
		cardBob.SortCurrentCard()
		cardBob.CheckCardLevel()
		cardAlice.SortCurrentCard()
		cardAlice.CheckCardLevel()
		matches.Matches[i].Result = CompareCard(&cardAlice.Card,&cardBob.Card)
	}
	end := time.Now().UTC().Nanosecond()
	fmt.Println(end-start)
	ReadFile("five_cards_with_ghost.result.json", &results)
	count := 0
	for i:=0; i < len(matches.Matches); i++ {
		if matches.Matches[i].Result != results.Matches[i].Result {
			count += 1
			fmt.Println(i,"result error", matches.Matches[i],results.Matches[i], count)
		}
	}
}

func TestCompareSevenCardWithGhostOneCore(t *testing.T)  {
	var results FiveCards
	var matches FiveCards
	ReadFile("seven_cards_with_ghost.json", &matches)
	start := time.Now().UTC().Nanosecond()
	for i := 0; i< len(matches.Matches); i++ {
		cardBob := CardSevenGhost{
			CardSeven:CardSeven{
				Card:Card{
					CurrentCard:matches.Matches[i].Bob,
				},
			},
		}
		cardAlice := CardSevenGhost{
			CardSeven:CardSeven{
				Card:Card{
					CurrentCard:matches.Matches[i].Alice,
				},
			},
		}
		cardBob.SortCurrentCard()
		cardBob.CheckCardLevel()
		cardAlice.SortCurrentCard()
		cardAlice.CheckCardLevel()
		matches.Matches[i].Result = CompareCard(&cardAlice.Card,&cardBob.Card)
	}
	end := time.Now().UTC().Nanosecond()
	fmt.Println(end-start)
	ReadFile("seven_cards_with_ghost.result.json", &results)
	count := 0
	for i:=0; i < len(matches.Matches); i++ {
		if matches.Matches[i].Result != results.Matches[i].Result {
			count += 1
			fmt.Println(i,"result error", matches.Matches[i],results.Matches[i], count)
		}
	}
}

func TestCompareSevenCardWithGhost(t *testing.T)  {
	var results FiveCards
	ReadFile("seven_cards_with_ghost.json", &matches)
	start := time.Now().UTC().Nanosecond()

	/*var wg sync.WaitGroup
	for i:=0;i< cores; i++ {
		wg.Add(1)
		go CompareCardWithFileSevenCardWithGhost(i,&wg)
	}
	wg.Wait()*/
	var chans = []chan int{}
	for i:=0;i < cores;i++ {
		c := make(chan int)
		chans = append(chans, c)
		go CompareCardWithFileSevenCardWithGhost(i,c)
	}
	for _,c := range chans {
		<- c
	}
	end := time.Now().UTC().Nanosecond()
	fmt.Println(end-start)
	ReadFile("seven_cards_with_ghost.result.json", &results)
	count := 0
	for i:=0; i < len(matches.Matches); i++ {
		if matches.Matches[i].Result != results.Matches[i].Result {
			count += 1
			fmt.Println(i,"result error", matches.Matches[i],results.Matches[i], count)
		}
	}
}

//func CompareCardWithFileSevenCardWithGhost(core int,wg *sync.WaitGroup)  {
func CompareCardWithFileSevenCardWithGhost(core int,c chan int)  {
	for i := 0; i< len(matches.Matches)/cores; i++ {
		cardBob := CardSevenGhost{
			CardSeven:CardSeven{
				Card:Card{
					CurrentCard:matches.Matches[i*cores+core].Bob,
				},
			},
		}
		cardAlice := CardSevenGhost{
			CardSeven:CardSeven{
				Card:Card{
					CurrentCard:matches.Matches[i*cores+core].Alice,
				},
			},
		}
		cardBob.SortCurrentCard()
		cardBob.CheckCardLevel()
		cardAlice.SortCurrentCard()
		cardAlice.CheckCardLevel()
		matches.Matches[i*cores+core].Result = CompareCard(&cardAlice.Card,&cardBob.Card)
	}
	//wg.Done()
	//fmt.Println(len(matches.Matches)/cores)
	c <- 1
}