package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"
)

func jsonFrom(source string) *interface{} {
	var file *os.File
	if source != "-" {
		var err error
		file, err = os.Open(source)
		if err != nil {
			log.Fatalf("Error reading data from: %s\n%s\n", source, err)
		}
		defer file.Close()
	} else {
		file = os.Stdin
	}
	decoder := json.NewDecoder(file)
	var data *interface{}
	err := decoder.Decode(&data)
	if err != nil {
		log.Fatalf("Error reading data from: %s\n%s\n", source, err)
	}
	return data
}

func printUsage() {
	fmt.Fprintln(os.Stderr, "Usage of jt: [OPTIONS] [FILE]")
	fmt.Fprintln(os.Stderr, "Renders JSON data read from FILE or stdin.\n")
	fmt.Fprintln(os.Stderr, "Options:")
	flag.PrintDefaults()
}

func printVersion() {
	fmt.Printf("jt v%s\n", VERSION)
	os.Exit(0)
}

func main() {
	templateFlag := flag.String("template", "", "(required) template file")
	versionFlag := flag.Bool("version", false, "print version")
	flag.Usage = printUsage
	flag.Parse()

	if *versionFlag {
		printVersion()
	}

	source := flag.Arg(0)
	if source == "" {
		source = "-"
	}
	if *templateFlag == "" {
		log.Fatalf("template is required")
	}

	t, err := template.ParseFiles(*templateFlag)
	if err != nil {
		log.Fatalf("Error loading template: %s\n%s\n", *templateFlag, err)
	}

	err = t.Execute(os.Stdout, jsonFrom(source))
	if err != nil {
		log.Fatalf("Error rendering template: %s\n", err)
	}
}
