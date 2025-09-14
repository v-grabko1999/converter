package webp

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"

	"rao/modules/converter"

	"golang.org/x/image/webp"
)

func init() {
	converter.RegisterDriver(&WebPDriver{})
}

// WebPDriver - драйвер для WebP формату
type WebPDriver struct{}

// ConvertToPNG конвертує байтовий масив WebP у байтовий масив PNG
func (d *WebPDriver) ConvertToPNG(input []byte) ([]byte, error) {
	// Створюємо новий рідер з байтового масиву
	imgReader := bytes.NewReader(input)

	// Декодуємо WebP зображення з рідера
	img, err := webp.Decode(imgReader)
	if err != nil {
		return nil, fmt.Errorf("error decoding WebP: %v", err)
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

func (d *WebPDriver) ConvertToJpg(input []byte) ([]byte, error) {
	// Створюємо новий рідер з байтового масиву
	imgReader := bytes.NewReader(input)

	// Декодуємо WebP зображення з рідера
	img, err := webp.Decode(imgReader)
	if err != nil {
		return nil, fmt.Errorf("error decoding WebP: %v", err)
	}

	var buf bytes.Buffer

	err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: 100})
	if err != nil {
		return nil, fmt.Errorf("error encoding PNG: %v", err)
	}

	return buf.Bytes(), nil
}

// Supports повертає список підтримуваних форматів
func (d *WebPDriver) Supports() []string {
	return []string{"webp"}
}
