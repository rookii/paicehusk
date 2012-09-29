package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/Rookii/paicehusk"
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

	for _, line := range testFile {
		words := reg.FindAllString(line, 1)
		for _, word := range words {
			fmt.Println(word, paicehusk.Stem(word, table))
		}
	}
}
