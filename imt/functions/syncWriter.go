package functions

import "io"

/**

Для синхронизации потока логгера и fmt

*/

type syncWriter struct {
	io.Writer
}

func (w syncWriter) Write(p []byte) (n int, err error) {
	return w.Writer.Write(p)
}
