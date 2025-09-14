package jpg

import (
	"bytes"
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/jpeg"
	"image/png"

	"rao/modules/converter"
)

func init() {
	converter.RegisterDriver(&JPGDriver{})
}

// JPGDriver - драйвер для JPG формату
type JPGDriver struct{}

// ConvertToPNG конвертує байтовий масив JPG у байтовий масив PNG
func (d *JPGDriver) ConvertToPNG(input []byte) ([]byte, error) {
	// Створюємо новий рідер з байтового масиву
	imgReader := bytes.NewReader(input)

	// Декодуємо JPG зображення з рідера
	img, err := jpeg.Decode(imgReader)
	if err != nil {
		return nil, fmt.Errorf("error decoding JPG: %v", err)
	}

	b := img.Bounds()
	palImg := image.NewPaletted(b, palette.Plan9)

	// Дизеринг (зменшує «смуги» кольорів)
	draw.FloydSteinberg.Draw(palImg, b, img, image.Point{})

	// Створюємо буфер для запису PNG
	var buf bytes.Buffer

	enc := png.Encoder{CompressionLevel: png.BestCompression}

	// Кодуючи зображення у PNG формат
	err = enc.Encode(&buf, palImg)
	if err != nil {
		return nil, fmt.Errorf("error encoding PNG: %v", err)
	}

	// Повертаємо результат як байтовий масив
	return buf.Bytes(), nil
}

func (d *JPGDriver) ConvertToJpg(input []byte) ([]byte, error) {
	// Повертаємо JPEG як []byte
	return input, nil
}

// Supports повертає список підтримуваних форматів
func (d *JPGDriver) Supports() []string {
	return []string{"jpg", "jpeg"}
}
