package main

import (
	"fmt"
	"math/rand"
	"os"
)

var (
	names = []string{
		"Jon",
		"Sheryl",
		"Andrew",
		"Gabbie",
		"Noah",
		"Tabby",
		"Jeremy",
		"Isaac",
		"Matthew",
	}
	restricted = map[string]string{
		"Jon":    "Sheryl",
		"Andrew": "Gabbie",
		"Noah":   "Tabby",
	}
)

func main() {
	assignments := make(map[string]string)

	var pool []string
	for _, n := range names {
		pool = append(pool, n)
	}

	for _, n := range names {
		var j int
		var n2 string
		for {
			for {
				j = rand.Intn(len(pool))
				n2 = pool[j]
				if n2 != n {
					break
				}
			}

			n2 = pool[j]
			if !(restricted[n] == n2 || restricted[n2] == n) {
				break
			}
		}

		assignments[n] = n2
		pool = append(pool[:j], pool[j+1:]...)
	}

	for k, v := range assignments {
		os.WriteFile("out/"+k+".html", []byte(v+"\n"), 0644)
	}

	links := ""
	for _, n := range names {
		links += fmt.Sprintf("<li><a href=\"%s.html\">%s</a></li>", n, n)
	}
	links = "<ul>" + links + "</ul>"

	index := "<h1>secret santa!</h1><p>Click your name to see your match... no snooping</p>" + links
	os.WriteFile("out/index.html", []byte(index), 0644)
}
