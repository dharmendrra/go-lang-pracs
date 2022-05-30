package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const chunkSize = 1000

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter string: ")
	s, _ := reader.ReadString('\n')

	log.Println("Reading prefixes list..")
	samplePrefixes, err := os.ReadFile("sample_prefixes.txt") //list of prefixes
	if err != nil {
		log.Fatal("Error opening Read File: ", err)
	}
	prefixes := strings.Split(string(samplePrefixes), "\n")

	p := prefixManager{}
	matchedPrefix := p.getLongestPrefixMatch(s, prefixes)
	if matchedPrefix != "" {
		log.Printf("Longest prefix match: %s \n", matchedPrefix)
	} else {
		log.Println("No prefix match found!")
	}
}
