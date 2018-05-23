package debug

import (
	"fmt"
	"testing"
	"time"

	DEBUG "github.com/computes/go-debug"
)

// ErrorFunction is a error debugging function
type ErrorFunction func(error, string, ...interface{})

// Debug creates a debug function for `name` which you call
// with printf-style arguments in your application or library.
func Debug(name string) DEBUG.DebugFunction {
	return DEBUG.Debug(name)
}

// DebugT creates a debug function for testing
// with printf-style arguments in your application or library.
func DebugT(t *testing.T) DEBUG.DebugFunction {
	var ts time.Time
	logged := false
	return func(format string, args ...interface{}) {
		message := fmt.Sprintf(format, args...)
		if logged {
			t.Logf("DEBUG: %s (%v)", message, time.Now().Sub(ts))
			ts = time.Now()
			return
		}
		t.Logf("DEBUG: %s", message)
		ts = time.Now()
		logged = true
	}
}

// Error creates a debug function for `name` with a `:error` suffix
// which you call with an error
// and printf-style arguments in your application or library.
// If err is nil, nothing will debugged
func Error(name string) ErrorFunction {
	debugError := DEBUG.Debug(fmt.Sprintf("%s:error", name))
	return func(err error, format string, args ...interface{}) {
		if err == nil {
			return
		}
		message := fmt.Sprintf(format, args...)
		debugError("%s: %+v", message, err)
	}
}

// ErrorT creates a debug function for testing
// which you call with an error
// and printf-style arguments in your application or library.
// If err is nil, nothing will debugged
func ErrorT(t *testing.T) ErrorFunction {
	return func(err error, format string, args ...interface{}) {
		if err == nil {
			return
		}
		message := fmt.Sprintf(format, args...)
		t.Errorf("%s: %+v", message, err)
	}
}
