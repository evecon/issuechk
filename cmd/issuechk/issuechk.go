/*
# Issue Checker

Scans your source code for issue links and prints there statuses.

## License

```
The MIT License (MIT)

Copyright (c) 2015 Evecon

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

```
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/evecon/issuechk/Godeps/_workspace/src/github.com/google/go-github/github"
	"os"
	"regexp"
	"strconv"
)

// https://github.com/evecon/issuechk/issues/1 https://github.com/evecon/issuechk/issues/2
func scanFile(client *github.Client, path string) {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	r, _ := regexp.Compile("https://github.com/([^/ ]+)/([^/ ]+)/issues/([0-9]+)")

	for ln := uint(0); scanner.Scan(); ln++ {
		for _, result := range r.FindAllStringSubmatch(scanner.Text(), -1) {
			checkIssue(client, path, ln, result[1], result[2], result[3])
		}
	}
}

func checkIssue(client *github.Client, path string, ln uint, org, repo, issueStr string) {
	issueNumber, perr := strconv.ParseInt(issueStr, 0, 32)
	if perr != nil {
		fmt.Printf("%s:%d: github.com/%s/%s issue #%s: error: %v\n",
			path, ln, org, repo, issueStr, perr)
	}

	issue, _, err := client.Issues.Get(org, repo, int(issueNumber))
	if err != nil {
		fmt.Printf("%s:%d: github.com/%s/%s issue #%s: error: %v\n",
			path, ln, org, repo, issueStr, err)
	} else {
		fmt.Printf("%s:%d: github.com/%s/%s issue #%s is %s\n",
			path, ln, org, repo, issueStr, *issue.State)
	}
}

var Version = "amnesiac"

func showVersionInfo() {
	fmt.Println("issuechk", Version)
}

var showVersion = false

func init() {
	const (
		showVersionUsage = "print version"
	)
	flag.BoolVar(&showVersion, "version", showVersion, showVersionUsage)
}

func main() {
	flag.Parse()
	if showVersion {
		showVersionInfo()
		return
	}

	path := flag.Arg(0)

	if len(path) > 0 {
		client := github.NewClient(nil)
		scanFile(client, path)
	}
}
