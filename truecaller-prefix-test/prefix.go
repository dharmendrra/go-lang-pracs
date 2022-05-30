package main

import (
	"strings"
	"sync"
)

type prefixManager struct {
}

func (p *prefixManager) getLongestPrefixMatch(s string, prefixes []string) string {
	var matchedPrefixList []string

	dispatcher := make(chan []string)
	result := make(chan string)

	p.dispatchPrefixes(dispatcher, result, s, prefixes)
	p.initializeWorkers(dispatcher, result, s)
	p.gatherMatchPrefixes(result, &matchedPrefixList)

	return p.getLongestPrefixFromList(matchedPrefixList)
}

func (p *prefixManager) getLongestPrefixFromList(prefixList []string) string {
	var matchedPrefix string
	// filter out longest prefix from the resulting matched prefixes
	for _, v := range prefixList {
		if len(v) >= len(matchedPrefix) {
			matchedPrefix = v
		}
	}
	return matchedPrefix
}

func (p *prefixManager) hasPrefix(s string, prefixList []string) string {
	var matchedPrefix string
	for _, v := range prefixList {
		if strings.HasPrefix(s, v) && len(v) >= len(matchedPrefix) {
			matchedPrefix = v
		}
	}
	return matchedPrefix
}

// dispatch chunks of prefixes in batches for concurrent matching
func (p *prefixManager) dispatchPrefixes(ch chan []string, result chan string, inputString string, prefixes []string) {
	go func() {
		defer close(ch)
		var endChunk int
		for startChunk := 0; startChunk < len(prefixes); startChunk = endChunk {
			endChunk = startChunk + chunkSize
			if endChunk > len(prefixes) {
				endChunk = len(prefixes)
			}
			// Create chunk of prefixes
			ch <- prefixes[startChunk:endChunk:endChunk]
		}
	}()
}

// initialize pool of routines
func (p *prefixManager) initializeWorkers(ch chan []string, result chan string, inputString string) {
	wg := new(sync.WaitGroup)
	num := 10 //number of routines in pool
	for i := 1; i <= num; i++ {
		wg.Add(1)
		go func(i int) {
			p.worker(ch, result, inputString)
			wg.Done()
		}(i)
	}
	go func() {
		wg.Wait()
		close(result)
	}()
}

// worker doing the matching from the chunk of prefixes
func (p *prefixManager) worker(ch chan []string, result chan string, inputString string) {
	for c := range ch {
		if s := p.hasPrefix(inputString, c); s != "" {
			result <- s
		}
	}
}

// gather longest matched prefixes from each chunk
func (p *prefixManager) gatherMatchPrefixes(result chan string, matchedPrefixList *[]string) {
	for r := range result {
		*matchedPrefixList = append(*matchedPrefixList, r)
	}
}
