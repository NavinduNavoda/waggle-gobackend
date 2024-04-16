package main

import (
	"github.com/NavinduNavoda/waggle-gobackend/data"
	"github.com/NavinduNavoda/waggle-gobackend/server"
)


func main() {

	dbinfo, err := data.ReadDBInfoFromFile()

	if err != nil {
		panic(err)
	}
	
	server := server.NewServer(dbinfo)
	
	server.Run(":5000")

}