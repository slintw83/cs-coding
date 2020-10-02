package main

import "fmt"

type Trie struct {
	IsWord   bool
	Children map[rune]*Trie
	Word     *string
}

func NewTrie() *Trie {
	trie := &Trie{Children: make(map[rune]*Trie)}
	return trie
}

func (trie *Trie) PopulateTrie(word string) {
	runes := []rune(word)
	node := trie
	for _, c := range runes {
		if _, ok := node.Children[c]; !ok {
			node.Children[c] = NewTrie()
		}
		node = node.Children[c]
	}
	node.IsWord = true
	node.Word = &word
}

func BoggleBoard(board [][]rune, words []string) []string {
	trie := NewTrie()
	for _, word := range words {
		trie.PopulateTrie(word)
	}
	fmt.Println(trie)

	found := make(map[string]bool)
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}

	for row := range board {
		for col := range board[row] {
			dfs(board, row, col, visited, found, *trie)
		}
	}

	result := make([]string, 0)
	for k := range found {
		result = append(result, k)
	}
	return result
}

func dfs(board [][]rune, row int, col int, visited [][]bool, found map[string]bool, trie Trie) {
	if row < 0 || row >= len(board) || col < 0 ||
		col < 0 || col >= len(board[0]) || visited[row][col] {
		return
	}

	currRune := board[row][col]
	if _, ok := trie.Children[currRune]; !ok {
		return
	}

	trie = *trie.Children[currRune]
	if trie.IsWord {
		_, ok := found[*trie.Word]
		if !ok {
			found[*trie.Word] = true
		}
	}

	visited[row][col] = true
	dfs(board, row-1, col, visited, found, trie)
	dfs(board, row-1, col-1, visited, found, trie)
	dfs(board, row, col-1, visited, found, trie)
	dfs(board, row+1, col-1, visited, found, trie)
	dfs(board, row+1, col, visited, found, trie)
	dfs(board, row+1, col+1, visited, found, trie)
	dfs(board, row, col+1, visited, found, trie)
	dfs(board, row-1, col+1, visited, found, trie)
	visited[row][col] = false
	return
}
