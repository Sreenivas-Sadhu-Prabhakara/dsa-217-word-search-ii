package main

import (
	"fmt"
	"sort"
)

type trieNode struct {
	children [26]*trieNode
	word     string // non-empty at a node that ends a word
}

// FindWords returns every dictionary word traceable in the grid through
// adjacent (4-directional) cells, each cell used once per word. One DFS per
// cell walks the board and the trie IN LOCKSTEP, so all words are hunted
// simultaneously; found words are cleared from the trie to prevent
// duplicates and shrink later searches.
func FindWords(board [][]byte, words []string) []string {
	root := &trieNode{}
	for _, w := range words {
		n := root
		for i := 0; i < len(w); i++ {
			c := w[i] - 'a'
			if n.children[c] == nil {
				n.children[c] = &trieNode{}
			}
			n = n.children[c]
		}
		n.word = w
	}

	out := []string{}
	rows := len(board)
	var dfs func(r, c int, n *trieNode)
	dfs = func(r, c int, n *trieNode) {
		if r < 0 || r >= rows || c < 0 || c >= len(board[r]) {
			return
		}
		ch := board[r][c]
		if ch == '#' {
			return // on the current path already
		}
		next := n.children[ch-'a']
		if next == nil {
			return // no dictionary word continues this way
		}
		if next.word != "" {
			out = append(out, next.word)
			next.word = "" // report once
		}
		board[r][c] = '#'
		dfs(r+1, c, next)
		dfs(r-1, c, next)
		dfs(r, c+1, next)
		dfs(r, c-1, next)
		board[r][c] = ch
	}
	for r := range board {
		for c := range board[r] {
			dfs(r, c, root)
		}
	}
	sort.Strings(out)
	return out
}

func main() {
	board := [][]byte{
		[]byte("oaan"),
		[]byte("etae"),
		[]byte("ihkr"),
		[]byte("iflv"),
	}
	fmt.Println(FindWords(board, []string{"oath", "pea", "eat", "rain"}))
}
