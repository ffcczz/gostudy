package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type FiveCard struct {
	Alice string `json:"alice"`
	Bob   string `json:"bob"`
	Result int `json:"result"`
}

type FiveCards struct {
	Matches []FiveCard `json:"matches"`
}

var matches FiveCards

func ReadFile(file string, matches *FiveCards)  {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		panic("wrong file")
	}
	err = json.Unmarshal(data, matches)
	if err != nil {
		fmt.Println(err)
	}
}

