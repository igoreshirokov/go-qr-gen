package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	content := flag.String("content", "https://itatarchenko.ru", "Данные для кодирования")
	output := flag.String("output", "qr.svg", "Имя выходного SVG файла")
	size := flag.Int("size", 200, "Размер изображения (квадрат)")
	flag.Parse()

	if *content == "" {
		fmt.Println("Ошибка: содержание QR-кода не может быть пустым")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("Создаём QR-код SVG: \"%s\"n", *content)
	err := generateQRCode(*content, *output, *size)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("QR-код сохранен в %s\n", output)
}

func generateQRCode(content, filename string, size int) error {
	qrCode, err := qr.Encode(content, qr.M,qr.Auto)
	if err != nil {
		return err
	}

	qrCode, err = barcode.Scale(qrCode, size, size)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return writeQRToSvg(file, qrCode, size)
}

func writeQRToSvg(file *os.File, code barcode.Barcode, size int) error {
	_, err := fmt.Fprintf(file, `<?xml version="1.0" encoding="UTF-8"?>
<svg width="%d" height="%d" viewBox="0 0 %d %d" xmlns="http://www.w3.org/2000/svg">
<rect width="100%%" height="100%%" fill="white"/>
`, size, size, size, size)

	width := code.Bounds().Dx()
	height := code.Bounds().Dy()
	blockSize := size / width

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// Check if pixel is dark (using color conversion)
			r, g, b, _ := code.At(x, y).RGBA()
			if r == 0 && g == 0 && b == 0 { // Black pixel
				_, err = fmt.Fprintf(file, `<rect x="%d" y="%d" width="%d" height="%d" fill="black"/>`,
					x*blockSize, y*blockSize, blockSize, blockSize)
				if err != nil {
					return err
				}
			}
		}
	}

	// SVG footer
	_, err = fmt.Fprintln(file, "</svg>")
	return err
}
