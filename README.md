[![GoDoc](https://godoc.org/github.com/flimzy/onload?status.png)](http://godoc.org/github.com/flimzy/onload)

**onload** is a very simple module, designed to provide cross-browser, jQuery-like onload
functionality without the bloat of jQuery.  It is writen entirely in Go.  To use it:

    package main
    
    import (
        "fmt"
        "github.com/flimzy/onload"
    )
    
    func main() {
        onload.Ready( func() {
            fmt.Print("The document is ready!\n")
        })
    }

This code is based on [this StackOverflow answer](http://stackoverflow.com/a/9899701/13860),
but somewhat simplified, and of course adapted to Go.

This software is released under the terms of the MIT license.  See LICENSE.txt for details.
