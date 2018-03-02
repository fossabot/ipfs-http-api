package dag

import (
	"runtime"

	DEBUG "github.com/computes/go-debug"
)

var debug = DEBUG.Debug("ipfs-http-api:dag")
var _debugStack = DEBUG.Debug("ipfs-http-api:dag:stack")

var debugStack = func() {
	stack := make([]byte, 1024)
	runtime.Stack(stack, false)

	_debugStack(string(stack))
}
