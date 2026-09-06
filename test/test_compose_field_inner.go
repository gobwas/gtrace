package test

//go:generate gtrace -v

//gtrace:gen
type InnerTrace struct {
	OnInner func()
}
