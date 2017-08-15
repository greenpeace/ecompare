package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

const emailRegex string = `([a-zA-Z0-9_\-\.]+)@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.)|(([a-zA-Z0-9\-]+\.)+))([a-zA-Z]{2,4}|[0-9]{1,3})(\]?)`

const shaRegex string = `[A-Fa-f0-9]{64}`

const dninieRegex string = `[A-z]?\d{7,8}[TRWAGMYFPDXBNJZSQVHLCKEtrwagmyfpdxbnjzsqvhlcke]`

var debug *bool

func main() {

	defer timeTrack(time.Now(), "main")

	help := flag.Bool("help", false, "Display help")
	data := flag.String("data", "emails", "What to compare")
	filenameA := flag.String("A", "a.txt", "File to do the operations")
	filenameB := flag.String("B", "b.txt", "File to do the operations")
	debug = flag.Bool("debug", false, "Debug the script")
	flag.Parse()

	if *help == true {
		helpMe()
	} else {

		// Read the files to strings
		aFile := fileToString(*filenameA)
		bFile := fileToString(*filenameB)

		// Create a map from the emails in each file with value as false
		var aMap, bMap map[string]bool
		switch *data {
		case "emails":
			aMap = searchInStringToMap(aFile, emailRegex)
			bMap = searchInStringToMap(bFile, emailRegex)
		case "sha256":
			aMap = searchInStringToMap(aFile, shaRegex)
			bMap = searchInStringToMap(bFile, shaRegex)
		case "dni":
			aMap = searchInStringToMap(aFile, dninieRegex)
			bMap = searchInStringToMap(bFile, dninieRegex)
		default:
			aMap = searchInStringToMap(aFile, emailRegex)
			bMap = searchInStringToMap(bFile, emailRegex)
		}

		// Transforms the values in the map to true when the key exits in the other map
		Compare(aMap, bMap)

		notInB := mapKeysToSlice(aMap, false)
		notInA := mapKeysToSlice(bMap, false)
		inBothAB := mapKeysToSlice(aMap, true)

		stringToFile("in-b-but-not-in-a.txt", strings.Join(notInA, "\n"))
		stringToFile("in-a-but-not-in-b.txt", strings.Join(notInB, "\n"))
		stringToFile("in-both-a-and-b.txt", strings.Join(inBothAB, "\n"))

		fmt.Println("\nWHAT HAPPENED?")
		fmt.Println("File A:", *filenameA)
		fmt.Println("File B:", *filenameB)
		fmt.Println("Parsed", *data, "in", *filenameA, ":", len(aMap))
		fmt.Println("Parsed", *data, "in", *filenameB, ":", len(bMap))
		fmt.Println("In", *filenameB, "but not in", *filenameA, ":", len(notInA), *data)
		fmt.Println("In", *filenameA, "but not in", *filenameB, ":", len(notInB), *data)
		fmt.Println("In both", *filenameA, "and", *filenameB, ":", len(inBothAB), *data)
		fmt.Println("Check the files:")
		fmt.Println("in-b-but-not-in-a.txt")
		fmt.Println("in-a-but-not-in-b.txt")
		fmt.Println("in-both-a-and-b.txt")
		fmt.Println("for more information.\n")
	}

}
