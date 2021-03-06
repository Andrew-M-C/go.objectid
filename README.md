# objectid

[![GoDoc](https://godoc.org/github.com/Andrew-M-C/go.objectid?status.svg)](https://godoc.org/github.com/Andrew-M-C/go.objectid)
[![Build Status](https://travis-ci.org/Andrew-M-C/go.objectid.svg?branch=dev-v1.0.3)](https://travis-ci.org/Andrew-M-C/go.objectid)
[![Coverage Status](https://coveralls.io/repos/github/Andrew-M-C/go.objectid/badge.svg)](https://coveralls.io/github/Andrew-M-C/go.objectid)
[![Stable](https://img.shields.io/badge/stable-v1.0.2-blue.svg)](https://github.com/Andrew-M-C/go.objectid/tree/v1.0.2)

Mongo object ID with 16-byte-length support (32 bytes in string), which could be replacement of UUID.

## What Is It for

MongoDB Object ID is a practical universal uniq ID solutions. But in some case, we need a 32-bytes long string to act like UUID, which also provides create time within. This what this package for.

The simplest way to implement 16-bytes long object ID (32 bytes as hex string) is adding random binaries in the tailing. However, I think nanoseconds are also quite randomized enough, thus I use nanoseconds instead.

## How to Use

### Create Object ID

To generate a 12-bytes standard object ID, Just use `New12`:

```go
id12 := objectid.New12()
```

It will return a `objectid.ObjectID`, equavilent to `[]byte`.

or you may use specified time:

```go
t := time.Now()
id12 := objectid.New12(t)
```

To generate a 16-bytes extended object ID, Just use `New16`, with the same usage of `New12`.

### Access Create Time

With `objectid.ObjectID` type, use `Time()` to get create time:

```go
id12 := objectid.New12()
id16 := objectid.New16()
tm12 := id12.Time()     // return time with accuracy of seconds
tm16 := id16.Time()     // return time with accuracy of nanoseconds, or microseconds in some OS
```

### Convert to Hexadecimal String

Use function `String()` to achieve this.

## Other Informations

[![License](https://img.shields.io/badge/license-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![](https://goreportcard.com/badge/github.com/Andrew-M-C/go.objectid)](https://goreportcard.com/report/github.com/Andrew-M-C/go.mysqlx)
[![codebeat badge](https://codebeat.co/badges/75922c43-e3af-4c63-8b8a-0e7da8f88acf)](https://codebeat.co/projects/github-com-andrew-m-c-go-objectid-master)
