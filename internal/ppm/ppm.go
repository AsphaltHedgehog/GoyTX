package ppm

import (
	"bufio"
	"fmt"
	"goytx/m/internal/vec"
	"io"
)

type Writer struct {
	w       *bufio.Writer
	width   int
	height  int
	written int
}

func NewWriter(out io.Writer, width, height int) (*Writer, error) {
	return &Writer{w: bufio.NewWriter(out), width: width, height: height}, nil
}

func (w *Writer) WritePixels(c vec.Color) error {
	r := c.X
	g := c.Y
	b := c.Z

	rbyte := int(255.999 * r)
	gbyte := int(255.999 * g)
	bbyte := int(255.999 * b)

	_, err := fmt.Fprintf(w.w, "%d %d %d\n", rbyte, gbyte, bbyte)

	if err != nil {
		return err
	}

	return nil
}

func (w *Writer) Close() error {
	if w.written != w.width*w.height {
		return fmt.Errorf("ppm: wrote %d pixels, header declares %d", w.written, w.width*w.height)
	}
	return w.w.Flush()
}
