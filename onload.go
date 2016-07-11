// +build js

// package onload provides a minimal, cross-browser document onload trigger
// for GopherJS. It is adapted from http://stackoverflow.com/a/9899701/13860
package onload

import (
	"github.com/gopherjs/gopherjs/js"
)

var readyList []func() = make([]func(), 0, 1)
var readyFired bool = false
var readyEventHandlersInstalled bool = false

// Ready registers a function to fire when the page is loaded. This is effectively
// the same as jQuery's $.ready(), but written in Go, and without the bloat of jQuery.
// If Ready() may be called multiple times, and once the document is ready, registered
// functions will be called in registration order.
func Ready(fn func()) {
	if readyFired {
		go fn()
		return
	}
	readyList = append(readyList, fn)
	doc := js.Global.Get("document")
	if doc.Get("readyState").String() == "complete" {
		go ready()
		return
	}
	if !readyEventHandlersInstalled {
		if doc.Get("addEventListener") != nil {
			// first choice is DOMContentLoaded event
			doc.Call("addEventListener", "DOMContentLoaded", ready, false)
			// backup is window load event
			js.Global.Call("addEventListener", "load", ready, false)
		} else {
			// Must be IE
			doc.Call("attachEvent", "onreadystatechange", readyStatechange)
			js.Global.Call("attachEvent", "onload", ready)
		}
		readyEventHandlersInstalled = true
	}
}

// ready does the actual work of running the registered functions when the
// document is finally ready
func ready() {
	if !readyFired {
		readyFired = true
		for _, readyFunc := range readyList {
			readyFunc()
		}
	}
}

// readyStateChange is a ready() wrapper for IE
func readyStatechange() {
	if js.Global.Get("document").Get("readyState").String() == "complete" {
		ready()
	}
}
