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
	bw := bufio.NewWriter(out)
	if _, err := fmt.Fprintf(bw, "P3\n%d %d\n255\n", width, height); err != nil {
		return nil, err
	}

	return &Writer{w: bw, width: width, height: height}, nil
}

func (w *Writer) WritePixels(c vec.Color) error {
	rbyte := int(255.999 * c.X)
	gbyte := int(255.999 * c.Y)
	bbyte := int(255.999 * c.Z)

	_, err := fmt.Fprintf(w.w, "%d %d %d\n", rbyte, gbyte, bbyte)

	if err != nil {
		return err
	}

	w.written++
	return nil
}

func (w *Writer) Close() error {
	if w.written != w.width*w.height {
		return fmt.Errorf("ppm: wrote %d pixels, header declares %d", w.written, w.width*w.height)
	}
	return w.w.Flush()
}
