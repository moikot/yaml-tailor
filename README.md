# yaml-tailor
![CI](https://github.com/moikot/yaml-tailor/workflows/CI/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/moikot/yaml-tailor)](https://goreportcard.com/report/github.com/moikot/yaml-tailor)

A tool for customizing YAML flies

## Run

```
Usage:
  yaml-tailor [file] [flags]

Flags:
  -h, --help                 help for yaml-tailor
  -s, --string stringArray   a string override
  -v, --value stringArray    a value override
```

The application uses [DJSON library syntax](https://github.com/moikot/djson#syntax) for interpreting string values and converting them to JSON objects and arrays. The same library
is used for [merging](https://github.com/moikot/djson#merging) the interpreted values with 
elements of an existing YAML file.

The examples below demonstrate how an empty YAML file is populated with `foo: bar` element. 

### Run as a standalone app

**Prerequisites:**
  * [Golang >=1.15](https://golang.org/doc/install)

```bash
$ go get github.com/moikot/yaml-tailor
$ echo '{}' > test.yml
$ yaml-tailor -v foo=bar test.yml
$ cat test.yml
```

### Run as a Docker container

**Prerequisites:**
  * [Docker](https://docs.docker.com/get-docker/)

```bash
$ echo '{}' > test.yml
$ docker run -d --rm -v $(pwd):/yml moikot/yaml-tailor -v foo=bar /yml/test.yml
$ cat test.yml
```
