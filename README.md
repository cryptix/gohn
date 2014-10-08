gohn
====
[![GoDoc](https://godoc.org/github.com/cryptix/gohn?status.svg)](https://godoc.org/github.com/cryptix/gohn)
[![Build Status](https://travis-ci.org/cryptix/gohn.svg?branch=master)](https://travis-ci.org/cryptix/gohn)

Golang package for the [new hacker news api](http://blog.ycombinator.com/hacker-news-api) at firebaseio.com.

It uses [bndr/gopencils](https://github.com/bndr/gopencils) for the http communication and json unmarshalling.


## TODO
- [] Users api
- [] Convert `time int` fields to `time time.Time`
- [] Add `TopStories()` which iterates over `TopStoryIDs()` to get the full storeis
