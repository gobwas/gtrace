package test

import "github.com/gobwas/gtrace/test/internal"

//go:generate gtrace -v

// NOTE: must compile without unused imports error.
//
//gtrace:gen
//gtrace:set context
type TraceNoShortcut struct {
	OnSomethingA func(Type)
	OnSomethingB func(internal.Type)
}
