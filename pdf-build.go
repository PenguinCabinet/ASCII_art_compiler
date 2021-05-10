package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"os"

	"github.com/goki/freetype/truetype"
	"github.com/signintech/gopdf"
)

func pdf_build(setting setting_file_t, source string) []byte {
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

	text_height := int(math.Sqrt(float64(setting.Font_size) * 72))

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: float64(imageWidth), H: float64(imageHeight)}})
	pdf.AddPage()

	err = pdf.AddTTFFont("font", "./font/font.ttf")
	if err != nil {
		panic(err)
	}

	pdf.SetFont("font", "", setting.Font_size)

	for i := 0; i < len(text); i++ {
		if len(text[i]) == 0 {
			continue
		}
		X := 0
		Y := ((i)*text_height + setting.Top_offset)

		pdf.SetX(float64(X))
		pdf.SetY(float64(Y))

		pdf.Cell(nil, text[i])
	}

	var A *bytes.Buffer = new(bytes.Buffer)

	pdf.Write(A)

	A_buf := A.Bytes()

	return A_buf
}
