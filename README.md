# turtle

A simple Golang web app that returns delayed results.

## Motivation

The intent of this application is that sometimes when you're writing tests,
you might have a test that timesout and doesn't wait long enough. To specifically
test these kinds of tests, this application is specifically slow so that you can
test timeout issues.

[I wrote a PHP web page specifically for this issue, but nice to have a native
 Golang version as well.]

## To Run

This is a golang application. You can run the app with:

```
$ PORT=8070 go run main.go
```

## To Do

- [x] Allow `port` to be specified in a 12-Factor way
- [x] Refactor code to conform to `golint` suggestions
- [x] Run `go fmt` to clean up code formatting
- [ ] Allow timeout to be specified
