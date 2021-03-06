// Code generated by gtrace. DO NOT EDIT.

package test

// Compose returns a new TraceReturningTrace which has functional fields composed
// both from t and x.
func (t TraceReturningTrace) Compose(x TraceReturningTrace) (ret TraceReturningTrace) {
	switch {
	case t.OnReturnedTrace == nil:
		ret.OnReturnedTrace = x.OnReturnedTrace
	case x.OnReturnedTrace == nil:
		ret.OnReturnedTrace = t.OnReturnedTrace
	default:
		h1 := t.OnReturnedTrace
		h2 := x.OnReturnedTrace
		ret.OnReturnedTrace = func() ReturnedTrace {
			r1 := h1()
			r2 := h2()
			switch {
			case r1.isZero():
				return r2
			case r2.isZero():
				return r1
			default:
				return r1.Compose(r2)
			}
		}
	}
	return ret
}
func (t TraceReturningTrace) onReturnedTrace() ReturnedTrace {
	fn := t.OnReturnedTrace
	if fn == nil {
		return ReturnedTrace{}
	}
	res := fn()
	return res
}
// Compose returns a new ReturnedTrace which has functional fields composed
// both from t and x.
func (t ReturnedTrace) Compose(x ReturnedTrace) (ret ReturnedTrace) {
	switch {
	case t.OnSomething == nil:
		ret.OnSomething = x.OnSomething
	case x.OnSomething == nil:
		ret.OnSomething = t.OnSomething
	default:
		h1 := t.OnSomething
		h2 := x.OnSomething
		ret.OnSomething = func(a int, b int) {
			h1(a, b)
			h2(a, b)
		}
	}
	switch {
	case t.OnFoo == nil:
		ret.OnFoo = x.OnFoo
	case x.OnFoo == nil:
		ret.OnFoo = t.OnFoo
	default:
		h1 := t.OnFoo
		h2 := x.OnFoo
		ret.OnFoo = func(i int, i1 int) {
			h1(i, i1)
			h2(i, i1)
		}
	}
	switch {
	case t.OnBar == nil:
		ret.OnBar = x.OnBar
	case x.OnBar == nil:
		ret.OnBar = t.OnBar
	default:
		h1 := t.OnBar
		h2 := x.OnBar
		ret.OnBar = func(i int, i1 int) {
			h1(i, i1)
			h2(i, i1)
		}
	}
	switch {
	case t.OnBaz == nil:
		ret.OnBaz = x.OnBaz
	case x.OnBaz == nil:
		ret.OnBaz = t.OnBaz
	default:
		h1 := t.OnBaz
		h2 := x.OnBaz
		ret.OnBaz = func(i int, i1 int) {
			h1(i, i1)
			h2(i, i1)
		}
	}
	return ret
}
// isZero checks whether t is empty
func (t ReturnedTrace) isZero() bool {
	if t.OnSomething != nil {
		return false
	}
	if t.OnFoo != nil {
		return false
	}
	if t.OnBar != nil {
		return false
	}
	if t.OnBaz != nil {
		return false
	}
	return true
}
func (t ReturnedTrace) onSomething(a int, b int) {
	fn := t.OnSomething
	if fn == nil {
		return
	}
	fn(a, b)
}
func (t ReturnedTrace) onFoo(i int, i1 int) {
	fn := t.OnFoo
	if fn == nil {
		return
	}
	fn(i, i1)
}
func (t ReturnedTrace) onBar(i int, i1 int) {
	fn := t.OnBar
	if fn == nil {
		return
	}
	fn(i, i1)
}
func (t ReturnedTrace) onBaz(i int, i1 int) {
	fn := t.OnBaz
	if fn == nil {
		return
	}
	fn(i, i1)
}
func traceReturningTraceOnReturnedTrace(t TraceReturningTrace) ReturnedTrace {
	res := t.onReturnedTrace()
	return res
}
