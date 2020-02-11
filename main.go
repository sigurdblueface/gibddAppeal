package main

import (
	"log"
	"os"
)

func main(){
	p := Processor{}
	err := p.Init()
	if err != nil{
		log.Printf("Error within initialization: %v", err)
	}
	err = p.Execute()
	if err != nil{
		log.Printf("Error while injecting values: %v", err)
		os.Exit(1)
	}
	os.Exit(0)
}