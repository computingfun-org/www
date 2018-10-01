package html

import (
	"bytes"
	"io"

	"github.com/shiyanhui/hero"
)

// NewBuffer returns a new buffer from hero's pool.
func NewBuffer() *bytes.Buffer {
	return hero.GetBuffer()
}

// CloseBuffer writers the buffer to the writer then returns the buffer back to hero's pool.
// If writer is nil the buffer is closed without being written.
func CloseBuffer(buffer *bytes.Buffer, w io.Writer) {
	if w != nil {
		w.Write(buffer.Bytes())
	}
	hero.PutBuffer(buffer)
}
