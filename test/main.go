package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/Rookii/paicehusk"
	"github.com/zaphar/go-stem/stemmer"
	"io"
	"os"
	"regexp"
)

func readFile(path string) (lines []string, err error) {
	var (
		file   *os.File
		part   []byte
		prefix bool
	)
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String(), "\n")
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	ruleFile, err := readFile("rules.txt")
	if err != nil {
		panic(err)
	}
	table := paicehusk.NewRuleTable(ruleFile)

	testFile, err := readFile("test.txt")
	if err != nil {
		panic(err)
	}
	reg := regexp.MustCompile("[\\w]+")

	count := 0

	for _, line := range testFile {
		//words := strings.Split(line, " ")
		words := reg.FindAllString(line, -1)
		for _, word := range words {
			porter := string(stemmer.Stem([]byte(word)))
			paice := paicehusk.Stem(word, table)
			if porter != paice {
				count++
				out.WriteString(word + " ")
				out.WriteString(porter + " ")
				out.WriteString(paice + "\n")
			}
		}
	}
	out.WriteString(fmt.Sprintf("%v", count))
}
