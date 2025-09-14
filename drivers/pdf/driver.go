package pdf

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"time"

	"github.com/v-grabko1999/converter"

	"github.com/google/uuid"
	"github.com/spf13/afero"
)

func init() {
	checkLibs()
	converter.RegisterDriver(&PDFDriver{fs: afero.NewMemMapFs()})
}

// PDFDriver - драйвер для PDF формату
type PDFDriver struct {
	fs afero.Fs // Віртуальна файлова система
}

// generateUniqueFileName генерує унікальне ім'я файлу на основі часу і UUID
func generateUniqueFileName(extension string) string {
	uniqueID := uuid.New().String()
	timeStamp := time.Now().Format("20060102_150405")
	return fmt.Sprintf("file_%s_%s.%s", timeStamp, uniqueID, extension)
}

// ConvertToPNG - спільна функція для роботи з різними ОС
func (d *PDFDriver) ConvertToPNG(input []byte) ([]byte, error) {
	// Створюємо унікальне ім'я файлу для зберігання у віртуальній файловій системі

	// Викликаємо специфічну для ОС функцію
	img, err := d.convert(input)
	if err != nil {
		return nil, err
	}

	// Створюємо буфер для запису PNG
	var buf bytes.Buffer

	// Конвертуємо зображення у PNG
	err = png.Encode(&buf, img)
	if err != nil {
		return nil, fmt.Errorf("error encoding PNG: %v", err)
	}

	return buf.Bytes(), nil
}

// ConvertToPNG - спільна функція для роботи з різними ОС
func (d *PDFDriver) ConvertToJpg(input []byte) ([]byte, error) {
	// Створюємо унікальне ім'я файлу для зберігання у віртуальній файловій системі

	// Викликаємо специфічну для ОС функцію
	img, err := d.convert(input)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: 100})
	if err != nil {
		return nil, fmt.Errorf("error encoding PNG: %v", err)
	}

	return buf.Bytes(), nil
}

// Supports повертає список підтримуваних форматів
func (d *PDFDriver) Supports() []string {
	return []string{"pdf"}
}
