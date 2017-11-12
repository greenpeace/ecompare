# Ecompare

**Compare two csv files to check which email addresses are in common and which email addresses are missing in each file.**

## What's the purpose of this script?

* We needed a quick-to-use **database debugging tool** for a vey specific problem: to help in a daily check about the information being sent from the mailing program to the CRM, and in the inverse direction.
* This script creates lists with the common and missing email addresses to be investigated. It can be useful for other purposes as well.
* Besides unique **email addresses** it can compare **sha256 hashes** or **Spanish ID numbers**.
* Parsing and comparing a 41MB file with 400.000 email addresses against another 36MB file with 350.000 email addresses takes 9 seconds in a fast laptop.

## How to use it

#### Get help

**ecompare** is a command line script. To get help type in the command line:

```bash
./ecompare --help
```

#### Export CSV files

You'll need to export csvs with supporter data in Engaging Networks and create similar exports in Salesforce.

#### Compare CSV files

Download both files as csvs to the same folder as the script. Then, using the command line run the script as:

```bash
./ecompare -data=emails -A=fileA.csv -B=fileB.csv
```

* `-data` specifies the type of data to compare. It can be `emails`, `sha256`, `urls` or `dni` (Spanish ID numbers)
* `-A` and `-B` specify the names of both files.

#### Get details about the comparison

When running the script you'll get a quick report like this:

```bash
WHAT HAPPENED?
File A: fileA.csv
File B: fileB.csv
Parsed emails in fileA.csv : 229
Parsed emails in fileB.csv: 214
In fileB.csv but not in fileA.csv : 0 emails
In fileA.csv but not in fileB.csv : 15 emails
In both fileA.csv and fileB.csv : 214 emails
```

And the script always creates 3 files in the current folder with the results:

1. *in-a-but-not-in-b.txt*
1. *in-b-but-not-in-a.txt*
1. *in-both-a-and-b.txt*

The filenames describe it's content and running the script again will overwrite this 3 files.

Open this files with your plain/code text editor and investigate the inconsistencies in both your CRM and mailing programs.

#### Delete files created by ecompare

To delete the 3 files created by ecompare:

```bash
./ecompare -trash
```

## Install

1. Download the [latest version of the binary code](https://github.com/greenpeace/gpes-ecompare/releases/) for your operating system to your desktop folder.
1. Unzip it to the desktop folder. *(Optionally copy the executable file to a folder in your [path](https://goo.gl/oLzTGw))*
1. To test your install, open the command line, go to the desktop folder and test it with the command:

* `./ecompare --help` *(Mac or Linux)*
* `./ecompare.exe --help` *(Windows)*

## Install from the source code

This script is also provided as [source code](https://github.com/greenpeace/gpes-ecompare/) in [Go](https://golang.org/dl/). To install:

```bash
go get github.com/greenpeace/gpes-ecompare
go install github.com/greenpeace/gpes-ecompare
```


## Note

This script works by parsing both text-based files with a regular expression rule to find emails, sha256 hashes or ID numbers. The advantage of using regular expressions to parse the files is speed, as the user doesn't has to adjust the files format. The disadvantage is that sometimes certain email addresses aren't grabbed. Mostly it will be invalid email addresses, but the script report will not match the CRM report.