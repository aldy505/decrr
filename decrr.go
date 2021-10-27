package decrr

import (
	"errors"
	"runtime"
	"strconv"
	"strings"
)

type frame struct {
	Func string
	Path string
	Line int
}

// Decorates a normal error and fill it with a stack trace.
// A modification of tracerr package.
func Wrap(err error) error {
	if err == nil {
		return nil
	}

	traced := trace(err)

	var stack strings.Builder
	for _, v := range traced {
		stack.WriteString("\n")
		stack.WriteString(v.Func)
		stack.WriteString(" ")
		stack.WriteString(v.Path)
		stack.WriteString(":")
		stack.WriteString(strconv.Itoa(v.Line))
	}
	return errors.New(err.Error() + "\n" + stack.String())
}

func trace(err error) []frame {
	skip := 2
	frames := make([]frame, 0, 20)
	for {
		pc, path, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		fn := runtime.FuncForPC(pc)
		frame := frame{
			Func: fn.Name(),
			Line: line,
			Path: path,
		}
		frames = append(frames, frame)
		skip++
	}

	return frames
}