# GoMiss
Auto `go get -u` from missing package message

## Installation

```
go get -u github.com/codenoid/GoMiss
```

## Usage

```
$ go build
a.go:9:2: cannot find package "github.com/gin-gonic/gin" in any of:
	/usr/local/go/src/github.com/gin-gonic/gin (from $GOROOT)
	/home/grandong/go/src/github.com/gin-gonic/gin (from $GOPATH)
a.go:10:2: cannot find package "github.com/codenoid/GoTral" in any of:
	/usr/local/go/src/github.com/codenoid/GoTral (from $GOROOT)
	/home/grandong/go/src/github.com/codenoid/GoTral (from $GOPATH)
a.go:11:2: cannot find package "github.com/getsentry/sentry-go" in any of:
	/usr/local/go/src/github.com/getsentry/sentry-go (from $GOROOT)
	/home/grandong/go/src/github.com/getsentry/sentry-go (from $GOPATH)
a.go:12:2: cannot find package "gopkg.in/couchbase/gocb.v1" in any of:
	/usr/local/go/src/gopkg.in/couchbase/gocb.v1 (from $GOROOT)
	/home/grandong/go/src/gopkg.in/couchbase/gocb.v1 (from $GOPATH)
a.go:13:2: cannot find package "github.com/go-redis/redis/v7" in any of:
	/usr/local/go/src/github.com/go-redis/redis/v7 (from $GOROOT)
	/home/grandong/go/src/github.com/go-redis/redis/v7 (from $GOPATH)
# When the app/package you run doesn't have go mod, then
# Really ? would you `go get` all these package manually ?
# here's GoMiss come to fix missing package for you
$ go build 2>&1 >/dev/null | GoMiss
$ go build
# success :)
```