+++
title = "Peeking into golang/x"
date = "2019-12-14T00:00:00+00:00"
series = ["Advent 2019"]
author = ["Ian Molee"]
+++

Everyone knows the Go standard library. It's readily presented, well documented,
and full of functionality. But what do you know about what lives in the shadowy
realm of golang.org/x?

This post will introduce the reader to `golang.org/x` and show a few examples
of useful packages that can be found there.

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

#### semaphore

### x/time

`x/time` contains a single package, `rate` , which implements a proactive rate
limiter. This is useful when you are calling a third party API whose rate limits
you know in advance.

E.g., 100 requests every 10 seconds.

Set up rate limiter for requests per second, and burst rate

### x/numbers

? Has to do with "message" package..?

Print numbers for locales/languages.
* Large numerals
* decimals
* Currency
* Percentage

### x/oauth2

### x/sys

Platform-specific stuff
??? Does this have stdlib equivalents?
???

### x/term

Determine if a program is running in a terminal.

### x/tools

#### x/tools/cmd/gorename

#### x/tools/cmd/eg

`eg` is a example-based refactoring tool. You provide it with an example of
code it should be on the lookout for to refactor, and how it should look after
refactoring. This can be useful for changing callsites across a codebase, for
instance, in a structural manner instead of just a textual manner.

--- fmt.Print/Printf example
--- changing call types example

## Conclusion

## About the Author

Ian Molee ([@ianfoo](https://twitter.com/ianfoo)) loves writing Go and talking
about Go. After being introduced to Go and being indoctrinated by its customs,
he could not believe he'd spent so much time writing things other than Go, and
he is fortunate to be able to program Go professionally.
