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

type Flags struct {
	template string
	version  bool
}

var flags = Flags{}

func init() {
	flag.StringVar(&flags.template, "template", "", "(required) template file")
	flag.BoolVar(&flags.version, "version", false, "print version")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, `
Usage: jt [OPTIONS] [FILE]
Renders JSON data read from FILE or stdin.

Options:`)
		flag.PrintDefaults()
	}

	flag.Parse()

	if flags.version {
		fmt.Printf("jt v%s\n", VERSION)
		os.Exit(0)
	}

	if flags.template == "" {
		log.Fatalf("template is required")
	}
}

func main() {
	source := flag.Arg(0)
	if source == "" {
		source = "-"
	}
	t, err := template.ParseFiles(flags.template)
	if err != nil {
		log.Fatalf("Error loading template: %s\n%s\n", flags.template, err)
	}

	err = t.Execute(os.Stdout, jsonFrom(source))
	if err != nil {
		log.Fatalf("Error rendering template: %s\n", err)
	}
}
