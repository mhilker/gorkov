package gorkov

import (
	"math/rand"
	"unicode/utf8"
)

type Chain struct {
	Graph map[string][]string
}

func NewChain() Chain {
	return Chain{make(map[string][]string)}
}

func (c *Chain) Train(parts []string) {
	for i, p := range parts {
		if i+1 < len(parts) {
 			c.Graph[p] = append(c.Graph[p], parts[i+1])
		}
	}
}

func (c *Chain) GenerateWithLength(length int) string {
	index := rand.Intn(len(c.Graph))
	var key string

	for key = range c.Graph {
		if index == 0 {
			break
		}
		index--
	}

	text := key

	for {
		l := len(c.Graph[key])
		if l == 0 {
			break
		}

		index = rand.Intn(l)
		key = c.Graph[key][index]
		text += key

		if utf8.RuneCountInString(text) >= length {
			break
		}
	}

	return text
}
