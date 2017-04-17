package writer

import (
	"fmt"
	"io"
)

// Writer is interface of Writer
type Writer interface {
	Print(a string)
	Printf(format string, a ...interface{})
}

// writer is strcut of writer
type writer struct {
	outStream io.Writer
}

// New creates Writer
func New(outStream io.Writer) Writer {
	return &writer{
		outStream: outStream,
	}
}

// Print formats using the default formats for its operands and writes to standard output.
func (w *writer) Print(a string) {
	fmt.Fprint(w.outStream, a)
}

// Printf formats according to a format specifier and writes to standard output.
func (w *writer) Printf(format string, a ...interface{}) {
	fmt.Fprint(w.outStream, fmt.Sprintf(format, a...))
}
