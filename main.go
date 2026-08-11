package main

import (
	"fmt"
	"io"
	"os"

	"github.com/AsphaltHedgehog/GoyTX/internal/geom"
	"github.com/AsphaltHedgehog/GoyTX/internal/ppm"
	"github.com/AsphaltHedgehog/GoyTX/internal/vec"
)

type Options struct {
	Seed     uint64
	Progress io.Writer
}

func main() {
	const imgWidth = 256
	const imgHeight = 256

	pw, err := ppm.NewWriter(os.Stdout, imgWidth, imgHeight)
	if err != nil {
		panic(err)
	}
	defer pw.Close()

	opts := Options{
		Seed:     0,
		Progress: os.Stderr,
	}

	for j := range imgHeight {
		fmt.Fprintf(opts.Progress, "\rScanlines remaining: %d ", imgHeight-j)
		for i := range imgWidth {
			pixelColor := vec.Color{
				float64(i) / float64(imgWidth-1),
				float64(j) / float64(imgHeight-1),
				0.0,
			}

			err := pw.WritePixels(pixelColor)

			if err != nil {
				panic(err)
			}

		}
	}

	if err := pw.Close(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	fmt.Fprint(os.Stderr, "\rDone.                 \n")
}
