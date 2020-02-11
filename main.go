package main

import (
	"log"
	"os"
)

func main(){
	p := (&Processor{}).Init()
	err := p.Execute()
	if err != nil{
		log.Printf("Error while injecting values: %v", err)
		os.Exit(1)
	}
	os.Exit(0)
}