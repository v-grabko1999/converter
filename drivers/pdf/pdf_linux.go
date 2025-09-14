//go:build linux
// +build linux

package pdf

import (
	"fmt"
	"image"
	"image/draw"

	"github.com/gen2brain/go-fitz"
)

// checkLibs перевіряє наявність бібліотеки go-fitz на Linux
func checkLibs() {

}

// convert - функція для Linux, яка конвертує PDF у PNG, повертає зображення
func (d *PDFDriver) convert(file []byte) (image.Image, error) {
	doc, err := fitz.NewFromMemory(file)
	if err != nil {
		return nil, fmt.Errorf("error opening PDF document: %v", err)
	}
	defer doc.Close()

	var imgs []image.Image
	totalHeight := 0
	maxWidth := 0

	// Конвертуємо кожну сторінку PDF у зображення
	for i := 0; i < doc.NumPage(); i++ {
		img, err := doc.Image(i)
		if err != nil {
			return nil, fmt.Errorf("error converting PDF page %d to image: %v", i+1, err)
		}
		imgs = append(imgs, img)

		totalHeight += img.Bounds().Dy()
		if img.Bounds().Dx() > maxWidth {
			maxWidth = img.Bounds().Dx()
		}
	}

	// Створюємо полотно для всіх сторінок
	finalImage := image.NewRGBA(image.Rect(0, 0, maxWidth, totalHeight))

	// Малюємо кожну сторінку на полотні
	currentY := 0
	for _, img := range imgs {
		draw.Draw(finalImage, image.Rect(0, currentY, img.Bounds().Dx(), currentY+img.Bounds().Dy()), img, image.Point{0, 0}, draw.Src)
		currentY += img.Bounds().Dy()
	}

	return finalImage, nil
}
