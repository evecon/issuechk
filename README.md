# Issue Checker

Scans your source code for issue links and prints there statuses.


## Use case

While working on a project you found a bug in a dependency
project. You submit an issue #XX to that project, make a workaround in
your project and annotate it with a comment:

```

    // This is a workaround to https://github.com/evecon/issuechk/issues/1
    ...workaround code goes here... 

```

Then you commit your code and continue development... Someday you grab
your project's code and want to check -- is that workaround still
needed or it can be properly fixed because that issue is closed. That
day you use *issuechk*.


## Download

Download a binary for you OS and platform from the [latest](https://github.com/evecon/issuechk/releases/latest) release.


## Quick start

To check one file, run:

```
issuechk program.c
```

## This release

- Search issues of the _GitHub_ hosted projects.
- Consume one issue link format `https://github.com/([^/ ]+)/([^/ ]+)/issues/([0-9]+)`.
- Output in format `file:line: github.com/{organization}/{repository} issue {issueNumber} is {state}` per issue.


## Build from sources

You will need a Go environment set up. Then run the following command:

```
go get github.com/evecon/issuechk/cmd/issuechk
```


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
