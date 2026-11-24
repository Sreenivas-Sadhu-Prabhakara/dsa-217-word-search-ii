package main

import (
	"reflect"
	"testing"
)

func grid(rows ...string) [][]byte {
	g := make([][]byte, len(rows))
	for i, r := range rows {
		g[i] = []byte(r)
	}
	return g
}

func TestFindWords(t *testing.T) {
	got := FindWords(grid("oaan", "etae", "ihkr", "iflv"), []string{"oath", "pea", "eat", "rain"})
	if !reflect.DeepEqual(got, []string{"eat", "oath"}) {
		t.Errorf("FindWords = %v, want [eat oath]", got)
	}

	// A cell may not be reused within one word.
	got = FindWords(grid("ab", "cd"), []string{"abcb"})
	if len(got) != 0 {
		t.Errorf("cell reuse accepted: %v", got)
	}

	// Each found word reported once even if traceable many ways.
	got = FindWords(grid("aa", "aa"), []string{"aaa"})
	if !reflect.DeepEqual(got, []string{"aaa"}) {
		t.Errorf("duplicates: %v", got)
	}

	// Words sharing prefixes are all found.
	got = FindWords(grid("ab"), []string{"a", "ab", "b"})
	if !reflect.DeepEqual(got, []string{"a", "ab", "b"}) {
		t.Errorf("shared prefixes: %v", got)
	}

	if got := FindWords(grid(), []string{"x"}); len(got) != 0 {
		t.Errorf("empty board: %v", got)
	}
}
