package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

var (
	Generator = NewNameGenerator()
)

type NameGenerator struct {
	Index int
	Names []string
}

func NewNameGenerator() *NameGenerator {
	return &NameGenerator{Index: 0, Names: generate()}
}

func (n *NameGenerator) Fetch() string {
	if n.Index == len(n.Names) {
		return ""
	}

	name := n.Names[n.Index]
	n.Index++

	return name
}

func generate() []string {
	const MAX = 26 * 26 * 10 * 10 * 10

	names := make([]string, MAX)
	index := 0

	for first := 'A'; first <= 'Z'; first++ {
		for second := 'A'; second <= 'Z'; second++ {
			for number := 0; number < 1000; number++ {
				names[index] = fmt.Sprintf("%c%c%03d", first, second, number)
				index++
			}
		}
	}

	rand.Shuffle(len(names), func(i, j int) { names[i], names[j] = names[j], names[i] })
	return names
}

// Robot has a unique name.
type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	r.name = Generator.Fetch()

	if r.name == "" {
		return "", errors.New("no names are available")
	}

	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}
