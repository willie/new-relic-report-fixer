package main

import (
	"bufio"
	"flag"
	//	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func processFile(in string, out string) {
	i, err := os.Open(in)
	if err != nil {
		log.Fatal(err)
	}
	defer i.Close()

	o, err := os.Create(out)
	if err != nil {
		log.Fatal(err)
	}
	defer o.Close()

	previousLine := ""

	scanner := bufio.NewScanner(i)
	for scanner.Scan() {
		line := scanner.Text()

		// don't care about End User Metrics -- they are borked from JS anyway
		if strings.Index(line, "End user") == 1 {
			continue
		}

		// don't want duplicate lines
		if line == previousLine {
			continue
		}

		o.WriteString(line + "\n")
		previousLine = line
	}
}

func main() {
	flag.Parse()

	for _, a := range flag.Args() {
		processFile(a, strings.TrimSuffix(a, filepath.Ext(a))+" fixed"+filepath.Ext(a))
	}
}
