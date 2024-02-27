package main

import (
	"encoding/json"
	"farmtrace"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alecthomas/jsonschema"
)

type schemaTuple struct {
	in       interface{}
	filename string
}

var schemaTuples = []schemaTuple{
	{farmtrace.DocTrace{}, "DocTrace-schema.json"},
	{farmtrace.Harvest{}, "Harvest-schema.json"},
	{farmtrace.DocLink{}, "DocLink-schema.json"},
	{farmtrace.Transfer{}, "Transfer-schema.json"},
}

func main() {
	fmt.Println(hdr)
	for _, st := range schemaTuples {
		err := schema2file(st.in, st.filename)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func schema2file(in interface{}, filename string) error {
	sch := genschema(in)
	fmt.Println(strings.Replace(schFl, "$FILENAME", filename, -1))
	return writeToFile(filename, sch)
}

func genSchema(in interface{}) string {
	schema := jsonschema.Reflect(in)
	json, err := json.MarshalIndent(schema, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

func writeToFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)

}

func genschema(in interface{}) string {
	schema := jsonschema.Reflect(in)
	json, err := json.MarshalIndent(schema, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

const (
	hdr = `
= Senior Care API JSON Schema

`
	schFl = `\n\n.$FILENAME\n[source,json]\n----\ninclude::$FILENAME[]\n\n----\n\n`
)
