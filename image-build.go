package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"math"
	"os"
	"strings"

	"github.com/goki/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func split_strings(text string) []string {
	//return strings.Split(text, "\n")
	text = strings.Replace(text, "\r\n", "\n", -1)
	return strings.FieldsFunc(text, func(c rune) bool {
		if c == '\n' || c == '\r' {
			return true
		} else {
			return false
		}
	})
}

func Get_img_wh(text []string, opt *truetype.Options, ft *truetype.Font, face font.Face, setting setting_file_t) (int, int) {

	//imageWidth := 10000
	//imageHeight := 10000

	//img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	/*
		dr := &font.Drawer{
			Dst:  img,
			Src:  image.Black,
			Face: face,
			Dot:  fixed.Point26_6{},
		}
	*/

	max_n := 0
	max_i := 0
	for i := 0; i < len(text); i++ {
		if max_n < len(text[i]) {
			max_n = len(text[i])
			max_i = i
		}
	}

	X := font.MeasureString(face, text[max_i]).Ceil()                                                //int(dr.MeasureString(text[max_i]))
	Y := len(text)*(face.Metrics().Ascent.Ceil()+face.Metrics().Descent.Ceil()) + setting.Top_offset //len(text) * int(dr.MeasureString("あ"))

	//fmt.Printf("%d\n", int(dr.MeasureString("あ")))

	return X, Y
}

func image_build(setting setting_file_t, source string) []byte {
	ftBinary, err := ioutil.ReadFile("font/font.ttf")
	ft, err := truetype.Parse(ftBinary)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	opt := truetype.Options{
		Size:              float64(setting.Font_size),
		DPI:               0,
		Hinting:           0,
		GlyphCacheEntries: 0,
		SubPixelsX:        0,
		SubPixelsY:        0,
	}

	face := truetype.NewFace(ft, &opt)

	text := split_strings(source)

	imageWidth, imageHeight := Get_img_wh(text, &opt, ft, face, setting)
	//fmt.Printf("Image size %d,%d\n", imageWidth, imageHeight)

	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	for x := 0; x < imageWidth; x++ {
		for y := 0; y < imageHeight; y++ {
			var c color.RGBA
			// それ以外は背景とする
			c = color.RGBA{255, 255, 255, 255}
			img.Set(x, y, c)
		}
	}

	text_height := int(math.Sqrt(float64(setting.Font_size)*72)) * int(face.Metrics().Ascent.Ceil()+face.Metrics().Descent.Ceil())

	dr := &font.Drawer{
		Dst:  img,
		Src:  image.Black,
		Face: face,
		Dot:  fixed.Point26_6{},
	}

	for i := 0; i < len(text); i++ {
		if len(text[i]) == 0 {
			continue
		}
		dr.Dot.X = 0
		//dr.Dot.X = (fixed.I(imageWidth) - dr.MeasureString(text[i])) / 2
		//dr.Dot.Y = fixed.Int26_6((imageHeight - len(text)*text_height + (i)*text_height))
		dr.Dot.Y = fixed.Int26_6((i)*text_height + imageWidth*2 + setting.Top_offset)
		dr.DrawString(text[i])
	}

	buf := &bytes.Buffer{}
	err = png.Encode(buf, img)

	buf_out := buf.Bytes()

	return buf_out
}
