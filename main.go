package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// 370103 words before using this
// same after?
func set(strings []string) []string {
	set := make(map[string]bool)
	for _, s := range strings {
		set[s] = true
	}
	result := make([]string, 0)
	for k := range set {
		result = append(result, k)
	}
	return result
}

// words reads from en.txt and returns a set of cleaned up words
func words() []string {
	content, err := ioutil.ReadFile("en.txt")
	if err != nil {
		log.Fatal(err)
	}
	return set(Map(strings.Split(string(content), "\n"), strings.TrimSpace))
}

func correct(word string) bool {
	list := words()
	for _, w := range list {
		if w == word {
			return true
		}
	}
	return false
}

func main() {
	word := os.Args[1]
	if correct(word) {
		fmt.Println("Correct")
	} else {
		fmt.Println("Incorrect")
	}
}
