package tree

import (
	"strings"

	l "github.com/mayukh42/goals/list"
	u "github.com/mayukh42/goals/utils"
)

type Trie struct {
	chars l.List
	term  bool
	Next  *Trie
}

func NewTrie() *Trie {
	l := make(l.List, 26)
	return &Trie{
		chars: l,
		term:  false,
		Next:  nil,
	}
}

func getIndex(c rune) int {
	return int(c) - 97
}

func (t *Trie) BuildTrie(word []rune) {
	if len(word) < 1 {
		t.term = true
		return
	}
	c := word[0]
	i := getIndex(c)
	t.chars[i] = 1
	t.Next = NewTrie()
	t.Next.BuildTrie(word[1:])
}

func (t *Trie) String() string {
	var sb strings.Builder
	sb.WriteRune('[')
	for i, x := range t.chars {
		if b, ok := x.(int); ok && b == 1 {
			c := rune(i + 97)
			sb.WriteRune(c)
		}
	}
	if t.term {
		sb.WriteRune('/')
	}
	sb.WriteRune(']')
	return sb.String()
}

func (t *Trie) ToSting(acc u.List) u.List {
	// postfix
	acc = append(acc, t.String())
	if t.Next != nil {
		return t.Next.ToSting(acc)
	}
	return acc
}
