package main

import "testing"

func TestHasPrefix(t *testing.T) {
	s := "UyhnnTBlApabesen" //input string
	prefixes := []string{"abc", "Uyhnn", "xsdfret", "random"}
	expectedMatch := "Uyhnn"
	p := &prefixManager{}
	match := p.hasPrefix(s, prefixes)
	if match != expectedMatch {
		t.Fatalf("Unexpected match. Want %s  Got %s", expectedMatch, match)
	}

}

func TestHasPrefixNotFound(t *testing.T) {
	s := "UyhnnTBlApabesen" //input string
	prefixes := []string{"abc", "iuUyhnn", "xsdfret", "random"}
	p := &prefixManager{}
	match := p.hasPrefix(s, prefixes)
	if match != "" {
		t.Fatalf("Unexpected match. Want %s  Got %s", "", match)
	}

}

func TestGetLongestPrefixFromList(t *testing.T) {
	matchPrefixes := []string{"abc", "iuUyhnn", "xsdfret", "random"}
	expectedOutput := "xsdfret"
	p := &prefixManager{}
	output := p.getLongestPrefixFromList(matchPrefixes)
	if output != expectedOutput {
		t.Fatalf("Unexpected output. Want %s  Got %s", expectedOutput, output)
	}
}

func TestGetLongestPrefixFromListNotFound(t *testing.T) {
	matchPrefixes := []string{}
	expectedOutput := ""
	p := &prefixManager{}
	output := p.getLongestPrefixFromList(matchPrefixes)
	if output != expectedOutput {
		t.Fatalf("Unexpected output. Want %s  Got %s", expectedOutput, output)
	}
}
