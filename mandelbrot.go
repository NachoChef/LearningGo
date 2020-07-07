package main

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
)

func main() {
	http.HandleFunc("/mandelbrot", CreateMandelbrot)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// CreateMandelbrot create and writes png mandelbrot img to client
func CreateMandelbrot(w http.ResponseWriter, r *http.Request) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	buffer := new(bytes.Buffer)
	png.Encode(buffer, img)
	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(buffer.Bytes())
}

func mandelbrot(z complex128) color.RGBA {
	const iterations = 200
	const contrast = 3
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			r, b := contrast*n, 200-contrast*n
			return color.RGBA{r, r - b, b, 255}
		}
	}
	return color.RGBA{}
}
