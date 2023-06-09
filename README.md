getgitignore
===

Tiny command line tool to create gitignore.

## Overview

Enter the development environment name or program language name in the argument,
the target environment `.gitignore` will be downloaded from [github](https://github.com/github/gitignore),
Create a `.gitignore` file in the current directory.
If a `.gitignore` file already exists, it will not be overwritten.

## Install

```
$ go install github.com/sharkattack51/getgitignore@latest
```

or [download release](https://github.com/sharkattack51/getgitignore/releases)

## Usage

```
$ [~/Desktop] mkdir test
$ [~/Desktop] cd ./test
$ [~/Desktop/test] getgitignore unity
getgitignore: [success] /Users/sharkattack51/Desktop/test/.gitignore
```

## Reference

https://github.com/github/gitignore
