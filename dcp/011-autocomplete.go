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

func (n *trieNode) getWords(f *[]string) {
	if n.word != nil {
		*f = append(*f, *n.word)
	}

	if (*n).children != nil {
		for _, v := range *n.children {
			(*v).getWords(f)
		}
	}
}

func (n *trieNode) query(q string) []string {
	found := []string{}
	r := []rune(q)
	curr := n

	for i := 0; i < len(r); i++ {
		if n, ok := (*curr.children)[r[i]]; ok {
			curr = n
		} else {
			return found
		}
	}

	(*curr).getWords(&found)
	return found
}

func (n *trieNode) insert(s string) {
	s = strings.TrimSpace(s)
	r := []rune(s)
	curr := n

	for i := 0; i < len(r); i++ {
		if _, ok := (*curr.children)[r[i]]; !ok {
			next := &trieNode{}
			(*next).initTrieRoot(r[i])
			(*curr.children)[r[i]] = next
			curr = next
		} else {
			curr = (*curr.children)[r[i]]
		}
	}

	(*curr).word = &s
}

func (n *trieNode) insertMany(input []string) {
	for _, s := range input {
		(*n).insert(s)
	}
}

func (n *trieNode) initTrieRoot(r rune) {
	children := make(map[rune]*trieNode)
	*n = trieNode{key: r, children: &children}
}
