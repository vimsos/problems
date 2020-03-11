package dcp

import "strings"

// Implement an autocomplete system. That is, given a query string s and a set of all possible query strings, return all strings in the set that have s as a prefix.
// For example, given the query string de and the set of strings [dog, deer, deal], return [deer, deal].
// Hint: Try preprocessing the dictionary into a more efficient data structure to speed up queries.

type trieNode struct {
	key      rune
	children *map[rune]*trieNode
	word     *string
}

var rootNode trieNode

func findWords(n *trieNode, f *[]string) {
	if n.word != nil {
		*f = append(*f, *n.word)
	}

	if (*n).children != nil {
		for _, v := range *n.children {
			findWords(v, f)
		}
	}
}

func query(q string) []string {
	found := []string{}
	r := []rune(q)
	curr := &rootNode

	for i := 0; i < len(r); i++ {
		if n, ok := (*curr.children)[r[i]]; ok {
			curr = n
		} else {
			return found
		}
	}

	findWords(curr, &found)
	return found
}

func insert(s string) {
	s = strings.TrimSpace(s)
	r := []rune(s)
	curr := &rootNode

	for i := 0; i < len(r); i++ {
		if _, ok := (*curr.children)[r[i]]; !ok {
			next := &trieNode{}
			initTrieNode(next, r[i])
			(*curr.children)[r[i]] = next
			curr = next
		} else {
			curr = (*curr.children)[r[i]]
		}
	}

	(*curr).word = &s
}

func insertMany(input []string) {
	for _, s := range input {
		insert(s)
	}
}

func initTrieNode(n *trieNode, r rune) {
	children := make(map[rune]*trieNode)
	*n = trieNode{key: r, children: &children}
}
