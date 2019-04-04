package main

import (
	"fmt"
	"github.com/go-interpreter/wagon/exec"
	"github.com/go-interpreter/wagon/wasm"
	"log"
	"os"
)

func main()  {
	//f, err := os.Open("/Users/konggan/workspace/go/src/gostudy/wasm/test.wasm")
	f, err := os.Open("/Users/konggan/workspace/go/src/gostudy/wasm/cmain.wasm")
	//f, err := file.Open("wasm_first.wasm")
	if err != nil {
		log.Fatal("Open file", err)
	}
	m,err := wasm.ReadModule(f, nil)
	if err != nil {
		log.Fatal("ReadModule",err)
	}
	for name, e := range m.Export.Entries {
		fmt.Println("name export   ",  e.Index, name)
		fmt.Println(len(m.Function.Types))
		/*fidx := m.Function.Types[int(e.Index)]
		ftype := m.Types.Entries[int(fidx)]
		fmt.Printf(" ftype %s: %#v: %v \n", name,e, ftype)*/
	}
	fmt.Println(m, "m")
	vm, err := exec.NewVM(m)
	if err != nil {
		log.Fatal("NewVM   ", err)
	}
	result, err := vm.ExecCode(0)
	if err != nil {
		log.Fatal("ExecCode", err)
	}
	fmt.Printf("result(%[1]T):%[1]v\n", result)
}
