+++
title = "Peeking into golang/x"
date = "2019-12-14T00:00:00+00:00"
series = ["Advent 2019"]
author = ["Ian Molee"]
+++

Everyone knows about the Go standard library, but do you know about `golang.org/x`?
It wouldn't be surprising if you didn't. There's not a ton of readily-available
information about it. That's not all that surprising, because `golang.org/x` isn't
something that can be documented easily, because it's not really anything other
than the beginning of an import path shared by a number of diverse Go packages,

Everyone knows the Go standard library. It's readily presented, well documented,
and full of functionality. But what do you know about what lives in the shadowy
realm of golang.org/x?

This post will introduce the reader to `golang.org/x` and show a few examples
of useful packages that can be found there.

Full list
https://godoc.org/-/subrepo

## What is golang.org/x?

This is kind of a trick question, because `golang.org/x` isn't anything,
really. It's not a git repository and it's not a Go package. Utlimately, it's
just a naming convention. It's just a common prefix for a number of packages'
import paths. 

`golang.org/x` is, therefore, not raelly a collection of any sort, but sort of
a designation. All of the follow statements apply to be what can be found with
a `golang.org/x` import path prefix.

* Go code that's useful but maybe not something that should be in the standard library
* Go code that's being considered for the standard library
* ? Go code that has been integrated to the standard library but that lives on in golang.org/x (context?, syncmap?)
* Command line tools
* Tools used by the Go team to develop and build Go itself (build, review)
* Tools for the Go internet presence (blog, website(?), proxies for forwarding)

!! Find all the packages under `golang.org/x` at https://build.golang.org !!

## Selected examples

With the diversity and wealth of code that can be found in `golang.org/x`, it would be
foolish to try to cover everything. In this section we'll look at some functionality
that may prove useful or that is otherwise intersting for curiosity's sake.

!!Alphabetize!!
### x/sync

Sync contains functionality useful for concurrent code. 

#### errgroup

Run a group of goroutines and watch for errors, canceling the coordinating
context if an error occurs. Sort of an error-aware sync.WaitGroup, because
it collects errors from the goroutines and cancels execution of all goroutines
if an error occurs.

#### semaphore

You can make a semaphore with a buffered channel, but this implementation is
more robust.

### x/time

`x/time` contains a single package, `rate` , which implements a proactive rate
limiter. This is useful when you are calling a third party API whose rate limits
you know in advance.

E.g., 100 requests every 10 seconds.

Set up rate limiter for requests per second, and burst rate. Can use `Wait` to just
wait until it's send to request, `Allow` to determine instantaneously whether a
request can be sent, or `Reserve` to reserve capacity for n units of work at a
given time, and the result will indicate how long the caller must wait before it
sending requests in order to remain within the rate limit.

https://play.golang.org/p/RpnGNYdE-MU

### x/numbers

? Has to do with "message" package..?

Print numbers for locales/languages.
* Large numerals
* decimals
* Currency
* Percentage

### x/oauth2

Lots of providers supported in sub packages:
Amazon, Google, Facebook, Twitch, Fitbit, GitHub, Gitlab, Bitbucket...
JWS, JWT

### x/sys

Platform-specific stuff
??? Does this have stdlib equivalents?
???

### x/image/font

### x/xerrors

Partially added to Go 1.13. Can use this package to get equivalent errors 
functionality in older versions of Go. Specifically, the `Is`, `As`, and `Unwrap`
functions, and `Errorf` which is like `fmt.Errorf` in Go 1.13+.

### x/term

Determine if a program is running in a terminal.

### x/lint/golint

The Go linter! Not sure why this isn't included in the main installation.

### x/tools

#### x/tools/cmd/gorename

#### x/tools/cmd/eg

`eg` is a example-based refactoring tool. You provide it with an example of
code it should be on the lookout for to refactor, and how it should look after
refactoring. This can be useful for changing callsites across a codebase, for
instance, in a structural manner instead of just a textual manner.

--- fmt.Print/Printf example
--- changing call types example

#### x/text/cmd/gotext

Use to extract text from Go source code and replace instances of fmt with message
printers--for localization!

## Conclusion

## About the Author

Ian Molee ([@ianfoo](https://twitter.com/ianfoo)) loves writing Go and talking
about Go. After being introduced to Go and being indoctrinated by its customs,
he could not believe he'd spent so much time writing things other than Go, and
he is fortunate to be able to program Go professionally.
