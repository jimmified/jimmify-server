package db

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
)

var links []string

func initLinks() {
	content, err := ioutil.ReadFile("resources/links.txt")
	if err != nil {
		log.Fatal(err)
	}
	links = strings.Split(string(content), "\n")
}

func RandomLinks() []string {
	// urls of chosen random links
	var r []string
	// size of this map used to check when 3 unique links have been selected
	m := make(map[int]bool)

	for len(m) < 3 {
		i := rand.Intn(len(links))
		if !m[i] && links[i] != "" {
			m[i] = true
			r = append(r, links[i])
		}
	}
	return r
}
