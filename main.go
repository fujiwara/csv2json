package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"flag"
	"io"
	"log"
	"os"
)

func main() {
	var in io.ReadCloser
	var noHeader bool
	flag.BoolVar(&noHeader, "n", false, "no header line")
	flag.Parse()
	if len(flag.Args()) >= 1 {
		f, err := os.Open(flag.Arg(0))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		in = f
	} else {
		in = os.Stdin
	}
	defer in.Close()
	r := csv.NewReader(in)
	records := [][]string{}
	maxColumns := 0
LINES:
	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break LINES
			} else if errors.Is(err, csv.ErrFieldCount) {
				// ignore
			} else {
				log.Println("invalid record:", record, err)
			}
		}
		records = append(records, record)
		if len(record) > maxColumns {
			maxColumns = len(record)
		}
	}
	if noHeader {
		dumpCSV(os.Stdout, records)
	} else {
		dumpCSVIncludesHeader(os.Stdout, records, maxColumns)
	}
}

func dumpCSV(w io.Writer, records [][]string) {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	for _, record := range records {
		if err := enc.Encode(record); err != nil {
			panic(err)
		}
	}
}

func dumpCSVIncludesHeader(w io.Writer, records [][]string, maxColumns int) {
	header := make([]string, 0, maxColumns)
	header = records[0]
	output := []map[string]string{}
	for _, record := range records[1:] {
		m := map[string]string{}
		for i, h := range header {
			if len(record) <= i {
				continue
			}
			m[h] = record[i]
		}
		output = append(output, m)
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	for _, m := range output {
		if err := enc.Encode(m); err != nil {
			panic(err)
		}
	}
}
