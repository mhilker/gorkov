package gorkov

import (
	"math/rand"
	"strings"
	"testing"
)

func TestHello(t *testing.T) {
	want := "abcabcabca"

	rand.Seed(0)
	c := NewChain()
	c.Train(strings.Split("a b c", " "))
	c.Train(strings.Split("b c a", " "))
	c.Train(strings.Split("c a b", " "))

	if got := c.GenerateWithLength(10); got != want {
		t.Errorf("c.GenerateWithLength(10) = %q, want %q", got, want)
	}
}
