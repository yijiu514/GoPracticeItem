package main

import (
	_ "GoPracticeItem/routers"
	"log"
	"net/http"
)

func main()  {

	log.Println("Running at port 7070...")
	err := http.ListenAndServe(":7070", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}


}