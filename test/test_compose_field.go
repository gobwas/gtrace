package test

//go:generate gtrace -v

//gtrace:gen
type OuterTrace struct {
	Inner   InnerTrace
	OnOuter func()
}
