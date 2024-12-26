package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func generateImage(w http.ResponseWriter, name, fontStyle, bgColor, textColor string, textSize int) {
	firstLetter := strings.ToUpper(string(name[0]))

	img1 := image.NewRGBA(image.Rect(0, 0, 200, 200))
	addRandomColorPixels(img1)

	img2 := image.NewRGBA(image.Rect(0, 0, 200, 200))
	addLetterImage(img2, firstLetter, parseHexColor(textColor), textSize)

	mergedImg := mergeImages(img1, img2)

	w.Header().Set("Content-Type", "image/png")
	if err := png.Encode(w, mergedImg); err != nil {
		http.Error(w, "Failed to encode image", http.StatusInternalServerError)
		log.Printf("Failed to encode image: %v", err)
	}
}

func addRandomColorPixels(img *image.RGBA) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	squareSize := img.Bounds().Dx() / 10 // 1% of the image size

	// Generate up to 3 random colors
	colors := make([]color.RGBA, 3)
	for i := range colors {
		colors[i] = color.RGBA{
			R: uint8(rnd.Intn(256)),
			G: uint8(rnd.Intn(256)),
			B: uint8(rnd.Intn(256)),
			A: 255,
		}
	}

	for y := 0; y < img.Bounds().Dy(); y += squareSize {
		for x := 0; x < img.Bounds().Dx(); x += squareSize {
			col := colors[rnd.Intn(len(colors))]
			draw.Draw(img, image.Rect(x, y, x+squareSize, y+squareSize), &image.Uniform{col}, image.Point{}, draw.Src)
		}
	}
}

func addLetterImage(img *image.RGBA, letter string, col color.Color, textSize int) {
	point := fixed.Point26_6{fixed.Int26_6(100 * 64), fixed.Int26_6(150 * 64)}
	face := basicfont.Face7x13
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: face,
		Dot:  point,
	}
	d.DrawString(letter)
}

func mergeImages(img1, img2 *image.RGBA) *image.RGBA {
	mergedImg := image.NewRGBA(img1.Bounds())
	draw.Draw(mergedImg, img1.Bounds(), img1, image.Point{}, draw.Src)
	draw.Draw(mergedImg, img2.Bounds(), img2, image.Point{}, draw.Over)
	return mergedImg
}

func parseHexColor(s string) color.RGBA {
	c := color.RGBA{A: 0xff}
	switch len(s) {
	case 7:
		c.R = hexToByte(s[1:3])
		c.G = hexToByte(s[3:5])
		c.B = hexToByte(s[5:7])
	case 4:
		c.R = hexToByte(string(s[1]) + string(s[1]))
		c.G = hexToByte(string(s[2]) + string(s[2]))
		c.B = hexToByte(string(s[3]) + string(s[3]))
	}
	return c
}

func hexToByte(h string) byte {
	var b byte
	for _, c := range h {
		b = b<<4 | hexDigit(c)
	}
	return b
}

func hexDigit(c rune) byte {
	switch {
	case '0' <= c && c <= '9':
		return byte(c - '0')
	case 'a' <= c && c <= 'f':
		return byte(c - 'a' + 10)
	case 'A' <= c && c <= 'F':
		return byte(c - 'A' + 10)
	}
	return 0
}
