package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"time"
)

// timeTrack is used to debug each function by measuring how long it takes to execute.
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	if *debug == true {
		log.Printf("%s took %s", name, elapsed)
	}
}

// fileToString Reads a file into a sting.
func fileToString(fileName string) string {
	defer timeTrack(time.Now(), "fileToString")
	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(dat)
}

// searchInStringToMap Reads a string and returns all lowercased matches in the regular expression as map keys
func searchInStringToMap(total string, expression string) map[string]bool {
	defer timeTrack(time.Now(), "searchInStringToMap")
	r, err := regexp.Compile(expression)
	if err != nil {
		panic(err)
	}
	allMatches := r.FindAllString(total, -1)
	a := make(map[string]bool)
	for _, v := range allMatches {
		a[strings.ToLower(v)] = false
	}
	return a
}

// Compare Compares 2 maps with words as what to search and boleans false value. Transforms in true when the key exists in the other map.
func Compare(a map[string]bool, b map[string]bool) (map[string]bool, map[string]bool) {
	defer timeTrack(time.Now(), "Compare")
	var y bool
	for key, _ := range a {
		_, y = b[key]
		if y == true {
			a[key] = true
		}
	}
	for key, _ := range b {
		_, y = a[key]
		if y == true {
			b[key] = true
		}
	}
	return a, b
}

// mapKeysToSlice Adds the keys with val to a slice
func mapKeysToSlice(m map[string]bool, val bool) []string {
	defer timeTrack(time.Now(), "mapKeysToSlice")
	var result []string
	for k, v := range m {
		if v == val {
			result = append(result, k)
		}
	}
	return result
}

// stringToFile writes a string to a file.
func stringToFile(fileName string, dat string) {
	defer timeTrack(time.Now(), "stringToFile")
	err := ioutil.WriteFile(fileName, []byte(dat), 0644)
	if err != nil {
		panic(err)
	}
}
