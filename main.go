package main

import (
	"flag"
	"io"
	"os"

	"encoding/json"

	"gopkg.in/yaml.v2"
)

type internal map[string]interface{}

func main() {

	var _internal = &internal{}

	var inputFile string
	var outputFile string

	var reader io.ReadCloser = os.Stdin
	var writer io.WriteCloser = os.Stdout

	defer func() {
		writer.Close()
		reader.Close()
	}()

	flag.StringVar(&inputFile, "input", "", "input file")
	flag.StringVar(&outputFile, "output", "", "output file")
	flag.Parse()

	if inputFile != "" {
		if in, e := os.OpenFile(inputFile, os.O_RDONLY, 0755); e != nil {
			panic(e)
		} else {
			reader = in
		}
	}

	if outputFile != "" {
		if out, e := os.OpenFile(outputFile, os.O_CREATE|os.O_RDWR, 0644); e == nil {
			writer = out
		}
	}

	if err := json.NewDecoder(reader).Decode(_internal); err != nil {
		panic(err)
	}

	if err := yaml.NewEncoder(writer).Encode(_internal); err != nil {
		panic(err)
	}
}
