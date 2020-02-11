package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

type Processor struct {
	templateFile string
	valuesFile string
	values map[string]interface{}
}

func (p *Processor) Init() Processor{
	templateFile := flag.String("template", "./template.gotmpl", "pass here a template file path")
	valuesFile := flag.String("values", "./values.yaml", "pass here a values file path")
	flag.Parse()
	p.templateFile = *templateFile
	p.valuesFile = *valuesFile
	return *p
}

func (p * Processor) Execute() error{
	tpl, err := template.ParseFiles(p.templateFile)
	if err != nil {
		fmt.Printf("Error while parsing values file: %v", err)
		return err
	}
	valBuffer, err := ioutil.ReadFile(p.valuesFile)
	if err != nil {
		fmt.Printf("Could not read values file: %v", err)
		return err
	}
	err = yaml.Unmarshal(valBuffer, &p.values)
	if err != nil {
		log.Printf("Could not unmarshall values file, check if yaml is valid: %v")
		return err
	}
	err = tpl.Execute(os.Stdout, p.values)
	if err != nil {
		log.Printf("Couldn't merge template and values: %v", err )
		return err
	}
	return nil
}