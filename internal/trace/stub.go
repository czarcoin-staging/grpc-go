// Package trace disables importing golang.org/x/net/trace to reduce binary size.
package trace

import (
	"context"
	"fmt"
)

func NewEventLog(family, title string) EventLog {
	return EventLog{}
}

type EventLog []byte // we need something that can be implicitly converted to nil

func (EventLog) Printf(format string, a ...interface{}) {}
func (EventLog) Errorf(format string, a ...interface{}) {}
func (EventLog) Finish()                                {}

type Trace []byte

func New(family, title string) Trace {
	return Trace{}
}
func NewContext(ctx context.Context, tr Trace) context.Context {
	return ctx
}
func FromContext(ctx context.Context) (Trace, bool) {
	return Trace{}, false
}

func (Trace) LazyLog(x fmt.Stringer, sensitive bool)     {}
func (Trace) LazyPrintf(format string, a ...interface{}) {}
func (Trace) SetError()                                  {}
func (Trace) SetRecycler(f func(interface{}))            {}
func (Trace) SetTraceInfo(traceID, spanID uint64)        {}
func (Trace) SetMaxEvents(m int)                         {}
func (Trace) Finish()                                    {}
