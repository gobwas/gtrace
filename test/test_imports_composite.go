package test

import "github.com/gobwas/gtrace/test/internal"

//go:generate gtrace -v

// NOTE: must compile without unused imports error.
//
//gtrace:gen
//gtrace:set shortcut
type TraceComposite struct {
	OnPointer func(*internal.Type)
	OnSlice   func([]internal.Type)
	OnArray   func([2]internal.Type)
	OnChan    func(chan internal.Type)
	OnMap     func(map[string]internal.Type)
	OnMapKey  func(map[internal.Type]string)
	OnNested  func([]map[internal.Type]*[]internal.Type)
}
