package caller

import (
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// Freme means a stack frame information.
type Frame struct {
	File      string // full path
	ShortFile string // base file name
	Line      int    // line number
	Func      string // package path + function name
	ShortFunc string // function name only
}

// Format returns a string representation of the Frame.
func (f Frame) Format() string {
	return f.File + ":" + strconv.Itoa(f.Line) + " " + f.ShortFile
}

// FormatShort returns a short string representation of the Frame.
func (f Frame) FormatShort() string {
	return f.ShortFile + ":" + strconv.Itoa(f.Line) + " " + f.ShortFunc
}

// Here returns the Frame information of the caller.
// skip=0 means the immediate caller of Here.
func Here(skip int) Frame {
	// 0: runtime.Callers, 1: Here, 2: caller of Here
	const internalSkip = 2

	var pcs [1]uintptr
	n := runtime.Callers(internalSkip+skip, pcs[:])
	if n == 0 {
		return Frame{File: "unknown", ShortFile: "unknown", Line: -1, Func: "unknown", ShortFunc: "unknown"}
	}

	frame, _ := runtime.CallersFrames(pcs[:n]).Next()

	fn := frame.Function
	shortFn := shortFuncName(fn)

	file := frame.File
	return Frame{
		File:      file,
		ShortFile: filepath.Base(file),
		Line:      frame.Line,
		Func:      fn,
		ShortFunc: shortFn,
	}
}

// "pkg/subpkg.Func" → "Func"
func shortFuncName(full string) string {
	if full == "" {
		return "unknown"
	}
	// Find last dot after last slash
	if i := strings.LastIndex(full, "."); i >= 0 && i+1 < len(full) {
		return full[i+1:]
	}
	return full
}
