package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	xj "goxml2json"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("x2j <xml file>")
		return
	}
	// xml is an io.Reader
	xb, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
		return
	}
	xml := strings.NewReader(string(xb))
	json, err := xj.Convert(xml)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(json.String())

}
