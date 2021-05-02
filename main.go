package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"

	"github.com/agnivade/levenshtein"
)

// Map applies the given function to each string in the slice
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// set removes duplicates from a slice of strings
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
	dictionary := words()
	for _, w := range dictionary {
		if w == word {
			return true
		}
	}
	return false
}

// minDistance returns the word that has the minimum Levenshtein distance from
// the given word, from the list of correct words
func minDistance(word string) string {
	dictionary := words()
	min := math.MaxInt32
	indexOfClosest := 0
	for i, w := range dictionary {
		if distance := levenshtein.ComputeDistance(word, w); distance < min {
			min = distance
			indexOfClosest = i
		}
	}
	return dictionary[indexOfClosest]
}

func main() {
	word := os.Args[1]
	if correct(word) {
		fmt.Println("Correct")
	} else {
		fmt.Printf("Did you mean %s?\n", minDistance(word))
	}
}
