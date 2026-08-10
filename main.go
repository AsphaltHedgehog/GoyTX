package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const imgWidth = 256
	const imgHeight = 256

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	fmt.Fprintf(w, "P3\n%d %d\n255\n", imgWidth, imgHeight)

	for j := 0; j < imgWidth; j++ {
		for i := 0; i < imgHeight; i++ {
			r := float64(i) / float64(imgWidth-1)
			g := float64(j) / float64(imgHeight-1)
			b := 0.0

			ir := int(255.999 * r)
			ig := int(255.999 * g)
			ib := int(255.999 * b)

			fmt.Fprintf(w, "%d %d %d\n", ir, ig, ib)
		}
	}
}
