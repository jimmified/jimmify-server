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
	var r []string
	for i := 0; i < 3; i++ {
		r = append(r, links[rand.Intn(len(links))])
	}
	return r
}
