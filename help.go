package main

import "fmt"

func helpMe() {

	textToPrint := `

* * * HELP * * *

Script to compare unique data from two text files, named A and B

- unique data includes emails or sha256
- text files include csv, txt, sql or html

Use the options as in this example:

./compare -data=emails -A=fileA.csv -B=fileB.csv

This script always creates 3 files in the current folder with the results. 

1 - in-a-but-not-in-b.txt
2 - in-b-but-not-in-a.txt
3 - in-both-a-and-b.txt

Each time the script runs it overwrites this 3 files.

---------------------

Comand line options:

-help				Display this help
-data=emails			What to compare in the files. It can be "emails", "sha256" or "dni". By default it compares emails.
-A=fileA.csv			File A name
-B=fileB.csv			File B name
-debug=true			Debug the script					
-trash				To delete the 3 files created by ecompare

`
	fmt.Printf(textToPrint)

}
