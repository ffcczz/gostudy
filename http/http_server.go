package main

import (
	"net/http"
	"log"
)

func Redirect(w http.ResponseWriter, r *http.Request)  {

}

func main()  {
	http.Handle("/index/",http.StripPrefix("/index/", http.FileServer(http.Dir("/home/qydev/eos/go"))))
	http.HandleFunc("/", Redirect)
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		log.Fatal("    ", err)
	}
}
