package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/igoreshirokov/go-qr-gen/qrcode"
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

	fmt.Printf("Создаём QR-код SVG: \"%s\"\n", *content)
	err := qrcode.GenerateQRCode(*content, *output, *size)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("QR-код сохранен в %s\n", *output)
}

