package lua

import (
	"fmt"
)

func argErrorf(state *State, argAt int, format string, args ...interface{}) {
	// TODO: stack analysis and debugging info if available.
	msg := fmt.Sprintf(format, args...)
	argError(state, argAt, msg)
}

func argError(state *State, argAt int, msg string) {
	// TODO: stack analysis and debugging info if available.
	state.Errorf("bad argument #%d (%s)", argAt, msg)
}

func intError(state *State, argAt int) {
	if isNumber(state.get(argAt)) {
		argError(state, argAt, "number has not ineger representation")
	}
	typeError(state, argAt, IntType)
}

func typeError(state *State, argAt int, want Type) {
	// TODO: stack analysis and debugging info if available.
	state.Errorf("%s expected @ %d, got %s", want, argAt, state.Value(argAt).Type())
}