package main

// 20feb24: Parameterized the comma to a saperator variable
// go run csv2json.go -in product.csv -out product.json

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	undef := "undefined"
	fmt.Printf("TAB CSV to JSON util\n")
	var outFile = flag.String("out", undef, "Output file, JSON file")
	var inFile = flag.String("in", undef, "Input file,tab TSV with first row being title")
	var sep = flag.String("sep", "comma", "Separator, default is comma")

	flag.Parse()

	if strings.EqualFold(undef, *outFile) || strings.EqualFold(undef, *inFile) {
		fmt.Printf("Please specify input and output file, with optional seperator. Example \ncsv2json -in coffee.csv -out catalog.json -sep tab\n")
		return
	}

	if strings.EqualFold("tab", *sep) {
		*sep = "\t"
	} else if strings.EqualFold("comma", *sep) {
		*sep = ","
	} else if strings.EqualFold("semicolon", *sep) {
		*sep = ";"
	} else if strings.EqualFold("pipe", *sep) {
		*sep = "|"
	} else if strings.EqualFold("colon", *sep) {
		*sep = ":"
	} else if strings.EqualFold("space", *sep) {
		*sep = " "
	}
	err := genjson(*inFile, *outFile, *sep)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Success generated\n")
}

func genjson(inFile, outFile, sep string) error {
	fileIn, err := os.Open(inFile)
	if err != nil {
		return err
	}
	defer fileIn.Close()

	fileOut, err := os.Create(outFile)
	if err != nil {
		return err
	}
	defer fileOut.Close()

	rd := bufio.NewReader(fileIn)
	wr := bufio.NewWriter(fileOut)
	i := 1
	_, err = wr.WriteString("[\n")
	if err != nil {
		return err
	}
	var hdr []string
	for {

		ln, err := rd.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}
		if err == io.EOF {
			break
		}
		ln = strings.TrimSpace(ln)
		if ln != "" {
			if i == 1 {
				hdr = strings.Split(ln, sep)
				for j := 0; j < len(hdr); j++ {
					hdr[j] = strings.Replace(hdr[j], " ", "", -1)
				}
			} else {
				if i > 2 {
					_, err = wr.WriteString(",\n")
					if err != nil {
						return err
					}
				}
				err = genJson(&hdr, ln, wr)
			}
			if err != nil {
				return err
			}
			i++

		}
	}
	_, err = wr.WriteString("]\n")
	wr.Flush()
	if err != nil {
		return err
	}
	fmt.Printf("Total %d lines processed from %s to create JSON file %s\n", i, inFile, outFile)
	return nil
}

func genJson(hdr *[]string, ln string, wr *bufio.Writer) error {
	row := strings.Split(ln, ",")

	colCnt := len(row)
	hdrCnt := len(*hdr)

	if colCnt != hdrCnt {
		fmt.Printf("mismatch hdr len = %d, col len = %d\n", hdrCnt, colCnt)
		if hdrCnt < colCnt {
			colCnt = hdrCnt
		}
	}
	_, err := fmt.Fprintf(wr, "{ ")
	if err != nil {
		return err
	}
	comma := ""
	for i, hx := range *hdr {
		_, err = fmt.Fprintf(wr, "%s\"%s\":\"%s\"", comma, hx, strings.TrimSpace(row[i]))
		if err != nil {
			return err
		}
		comma = ", "
	}
	_, err = fmt.Fprintf(wr, "} ")
	if err != nil {
		return err
	}
	return nil
}
