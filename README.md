gocommon
========

Common Go library for Joyent's Triton and Manta.

## Installation

Use `go-get` to install gocommon.
```
go get github.com/joyent/gocommon
```

## Documentation

Auto-generated documentation can be found on godoc.

- [github.com/joyent/gocommon](http://godoc.org/github.com/joyent/gocommon)
- [github.com/joyent/gocommon/client](http://godoc.org/github.com/joyent/client)
- [github.com/joyent/gocommon/errors](http://godoc.org/github.com/joyent/gocommon/errors)
- [github.com/joyent/gocommon/http](http://godoc.org/github.com/joyent/gocommon/http)
- [github.com/joyent/gocommon/jpc](http://godoc.org/github.com/joyent/gocommon/jpc)
- [github.com/joyent/gocommon/testing](http://godoc.org/github.com/joyent/gocommon/testing)


## Contributing

Report bugs and request features using [GitHub Issues](https://github.com/joyent/gocommon/issues), or contribute code via a [GitHub Pull Request](https://github.com/joyent/gocommon/pulls). Changes will be code reviewed before merging. In the near future, automated tests will be run, but in the meantime please `go fmt`, `go lint`, and test all contributions.


## Developing

This library assumes a Go development environment setup based on [How to Write Go Code](https://golang.org/doc/code.html). Your GOPATH environment variable should be pointed at your workspace directory.

You can now use `go get github.com/joyent/gocommon` to install the repository to the correct location, but if you are intending on contributing back a change you may want to consider cloning the repository via git yourself. This way you can have a single source tree for all Joyent Go projects with each repo having two remotes -- your own fork on GitHub and the upstream origin.

For example if your GOPATH is `~/src/joyent/go` and you're working on multiple repos then that directory tree might look like:

```
~/src/joyent/go/
|_ pkg/
|_ src/
   |_ github.com
      |_ joyent
         |_ gocommon
         |_ gomanta
         |_ gosdc
         |_ gosign
```

### Recommended Setup

```
$ mkdir -p ${GOPATH}/src/github.com/joyent
$ cd ${GOPATH}/src/github.com/joyent
$ git clone git@github.com:<yourname>/gocommon.git

# fetch dependencies
$ git clone git@github.com:<yourname>/gosign.git
$ go get -v -t ./...

# add upstream remote
$ cd gocommon
$ git remote add upstream git@github.com:joyent/gocommon.git
$ git remote -v
origin  git@github.com:<yourname>/gocommon.git (fetch)
origin  git@github.com:<yourname>/gocommon.git (push)
upstream        git@github.com:joyent/gocommon.git (fetch)
upstream        git@github.com:joyent/gocommon.git (push)
```

### Run Tests

```
cd ${GOPATH}/src/github.com/joyent/gocommon
go test ./...
```

### Build the Library

```
cd ${GOPATH}/src/github.com/joyent/gocommon
go build ./...
```
