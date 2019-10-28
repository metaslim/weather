# Simple Weather code

Access using browser at http://localhost:8080/weather?city=melbourne

## Caveat
The key is harcoded in config now.

## Environment

Go 1.13

## Setting up locally

### GOPATH and installing the code
```
$ mkdir -p $HOME/Development/gocode
$ export GOPATH=$HOME/Development/gocode
$ go get -v github.com/metaslim/weather
$ cd $GOPATH/src/github.com/metaslim/weather
$ git status
On branch master
Your branch is up to date with 'origin/master'.

nothing to commit, working tree clean
```

### Install Dependencies
```sh
$ cd v1
$ make install
```

## How to run tests

```sh
$ cd v1
$ make test

```

## How to build binary

```sh
$ cd v1
$ make build
```

## How to run binary

```sh
$ cd v1
$ ./weather
```

## How to run without compiling (slower compared to run binary)

```sh
$ cd v1
$ make run
```
