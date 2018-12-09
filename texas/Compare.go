package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FiveCard struct {
	Alice string `json:"alice"`
	Bob   string `json:"bob"`
	Result string `json:"result"`
}

type FiveCards struct {
	Matches []FiveCard `json:"matches"`
}

var matches FiveCards

func ReadFile(file string)  {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		panic("wrong file")
	}
	json.Unmarshal(data, &matches)
}


func main()  {
	currentPath,_ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(currentPath)
	ReadFile("/private/var/folders/tt/8v4bb03x685c49jk_vkvr7s80000gp/T/match.go")
	fmt.Println(matches)
}
