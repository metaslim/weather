# Simple Weather code

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
$ make install
```

## How to run tests

```sh
$ make test

```

## How to build binary

```sh
$ make build
```

## How to run binary

```sh
$ ./main
```

## How to run without compiling (slower compared to run binary)

```sh
$ make run
```
