# README #

![Alt text](http://natebrennand.github.io/concurrency_and_golang/pics/gopher_head.png)

## Overview
This is a standalone command line based, lightweight tool that will process flat files in a predefined format and will generate a JSON file into an indicated path

## Stack choice and architecture
[Golang](https://golang.org/) due to idiomatic baked in features (JSON processing, buffered read/write, etc). Performance gains over Java due to leaner abstraction over CPU

## Install go
Instructions [here](https://golang.org/doc/install)
* Set up your `.profile`
    export GOROOT=$HOME/go
    export GOPATH=$HOME/work
* Use `go get bitbucket.org/camilo_crespo/camilocrespo`, alternatively create `src/bitbucket.org/camilo_crespo/camilocrespo` folder in `GOPATH` and clone this repo there

## Run the app
    go run transformer.go input_path output_path

## Run the tests
    go test -v

## Generate documentation
    godoc -http=:6060






