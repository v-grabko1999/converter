package bmp

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"

	"golang.org/x/image/bmp"

	"github.com/v-grabko1999/converter"
)

func init() {
	converter.RegisterDriver(&BMPDriver{})
}

// BMPDriver - драйвер для BMP формату
type BMPDriver struct{}

// ConvertToPNG конвертує байтовий масив BMP у байтовий масив PNG
func (d *BMPDriver) ConvertToPNG(input []byte) ([]byte, error) {
	// Створюємо новий рідер з байтового масиву
	imgReader := bytes.NewReader(input)

	// Декодуємо BMP зображення з рідера
	img, err := bmp.Decode(imgReader)
	if err != nil {
		return nil, fmt.Errorf("error decoding BMP: %v", err)
	}

	// Створюємо буфер для запису PNG
	var buf bytes.Buffer

	// Кодуючи зображення у PNG формат
	err = png.Encode(&buf, img)
	if err != nil {
		return nil, fmt.Errorf("error encoding PNG: %v", err)
	}

	// Повертаємо результат як байтовий масив
	return buf.Bytes(), nil
}

func (d *BMPDriver) ConvertToJpg(input []byte) ([]byte, error) {
	// Створюємо reader з BMP-даних
	imgReader := bytes.NewReader(input)

	// Декодуємо BMP-зображення
	img, err := bmp.Decode(imgReader)
	if err != nil {
		return nil, fmt.Errorf("error decoding BMP: %v", err)
	}

	// Підготуємо буфер для запису JPG
	var buf bytes.Buffer

	opts := &jpeg.Options{Quality: 100}
	if err := jpeg.Encode(&buf, img, opts); err != nil {
		return nil, fmt.Errorf("error encoding JPG: %v", err)
	}

	// Повертаємо JPEG як []byte
	return buf.Bytes(), nil
}

// Supports повертає список підтримуваних форматів
func (d *BMPDriver) Supports() []string {
	return []string{"bmp"}
}
