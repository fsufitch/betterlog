package betterlog

import "runtime"

// CallerLoc is a container struct for data returned from [runtime.Caller]
type CallerLoc struct {
	ProgramCounter uintptr
	Filename       string
	LineNumber     int
	OK             bool
}

// GetCaller is a wrapper around [runtime.Caller] which encapsulates the results in a [CallerLoc] struct.
// If OK is flase, then Filename == "" and LineNumber == -1.
// Setting "skip" to 2 returns the caller of the calling function (same as in [runtime.Caller]).
func GetCaller(skip int) (c CallerLoc) {
	c.ProgramCounter, c.Filename, c.LineNumber, c.OK = runtime.Caller(skip + 1)
	if !c.OK {
		c.Filename = ""
		c.LineNumber = -1
	}
	return
}

type Traceback []CallerLoc

const DefaultMaxTracebackDepth = 32

// GetTraceback returns traceback of the callers leading to the current execution point.
// The returned list of callers is in reverse order (the caller of [GetTraceback] is first).
// "offset" indicates how many callers should be skipped (0 means none).
// "max" indicates the maximum depth to follow the stack to; 0 will default to [DefaultMaxTracebackDepth]; values below 0 will remove the upper limit (set with care).
// If the traceback overflows "max", then "overflow" will be set true.
func GetTraceback(offset int, max int) (tb Traceback, overflow bool) {
	if max == 0 {
		max = DefaultMaxTracebackDepth
	}

	skip := offset + 1
	for {
		current := GetCaller(skip + 1) // +1 skips GetTraceback itself
		if !current.OK {
			return
		}
		if max >= 1 && len(tb) >= max {
			// Trace extends beyond the set max
			overflow = true
			return
		}
		tb = append(tb, current)
		skip++
	}
}
