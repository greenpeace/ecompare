/*
.

Compare two csv files to check which email addresses are in common and which email addresses are missing in each file.

WHAT'S THE PURPOSE OF THIS SCRIPT?

- We needed a quick-to-use database debugging tool for a vey specific problem: to help in a daily check about the information being sent from the mailing program to the CRM, and in the inverse direction.
- This script creates lists with the common and missing email addresses to be investigated. It can be useful for other purposes as well.
- Besides unique email addresses it can compare sha256 hashes or Spanish ID numbers.
- Parsing and comparing a 41MB file with 400.000 email addresses against another 36MB file with 350.000 email addresses takes 9 seconds in a fast laptop.


GET HELP

./ecounter --help

.
*/
package main
