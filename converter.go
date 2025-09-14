package converter

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"runtime"
)

// Реєстратор для зберігання драйверів
var drivers []Driver

// RegisterDriver реєструє новий драйвер, але якщо драйвер вже зареєстрований, виводить попередження
func RegisterDriver(driver Driver) {
	// Перевіряємо, чи вже є драйвер, що підтримує ті ж самі формати
	for _, d := range drivers {
		if hasDuplicateFormats(d.Supports(), driver.Supports()) {
			// Якщо драйвер вже зареєстрований, виводимо попередження і стек викликів
			fmt.Println("Попередження: цей драйвер вже зареєстрований!")
			printStack()
			return
		}
	}
	// Якщо драйвер ще не зареєстрований, додаємо його
	drivers = append(drivers, driver)
}

// hasDuplicateFormats перевіряє наявність однакових форматів у двох списках
func hasDuplicateFormats(formats1, formats2 []string) bool {
	formatSet := make(map[string]struct{})
	for _, format := range formats1 {
		formatSet[format] = struct{}{}
	}
	for _, format := range formats2 {
		if _, exists := formatSet[format]; exists {
			return true
		}
	}
	return false
}

// Convert - конвертує файл у PNG, шукаючи відповідний драйвер
func Convert(input []byte, format string) ([]byte, error) {
	for _, driver := range drivers {
		for _, supportedFormat := range driver.Supports() {
			if supportedFormat == format {
				return driver.ConvertToPNG(input)
			}
		}
	}
	return nil, errors.New("unsupported format")
}

func ConvertToJpg(input []byte, format string) ([]byte, error) {
	for _, driver := range drivers {
		for _, supportedFormat := range driver.Supports() {
			if supportedFormat == format {
				return driver.ConvertToJpg(input)
			}
		}
	}
	return nil, errors.New("unsupported format")
}

func ConvertAuto(input []byte) ([]byte, error) {
	dt := DetectFormat(input)
	if dt == "png" {
		return input, nil
	}
	for _, driver := range drivers {
		for _, supportedFormat := range driver.Supports() {
			if supportedFormat == dt {
				return driver.ConvertToPNG(input)
			}
		}
	}
	return nil, errors.New("unsupported format: " + dt)
}

func ConvertAutoToJpg(input []byte) ([]byte, error) {
	dt := DetectFormat(input)
	if dt == "png" {
		return input, nil
	}
	for _, driver := range drivers {
		for _, supportedFormat := range driver.Supports() {
			if supportedFormat == dt {
				return driver.ConvertToJpg(input)
			}
		}
	}
	return nil, errors.New("unsupported format: " + dt)
}

// printStack виводить стек викликів
func printStack() {
	buf := make([]byte, 1024)
	n := runtime.Stack(buf, false)
	fmt.Printf("%s\n", buf[:n])
}

func DetectFormat(data []byte) string {
	if len(data) > 8 && bytes.Equal(data[:8], []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}) {
		return "png"
	}

	// Проверяем BMP (начинается с "BM")
	if len(data) > 2 && bytes.Equal(data[:2], []byte{0x42, 0x4D}) {
		return "bmp"
	}

	// Проверяем JPEG (начинается с 0xFFD8 и заканчивается на 0xFFD9)
	if len(data) > 2 && bytes.Equal(data[:2], []byte{0xFF, 0xD8}) {
		if len(data) > 2 && bytes.Equal(data[len(data)-2:], []byte{0xFF, 0xD9}) {
			return "jpeg"
		}
	}

	// Проверяем PDF (начинается с "%PDF-")
	if len(data) > 5 && bytes.Equal(data[:5], []byte("%PDF-")) {
		return "pdf"
	}

	// Проверяем WebP (начинается с "RIFF....WEBP")
	if len(data) > 12 && bytes.Equal(data[:4], []byte("RIFF")) && bytes.Equal(data[8:12], []byte("WEBP")) {
		return "webp"
	}

	switch http.DetectContentType(data[:min(len(data), 512)]) {
	case "image/jpeg":
		return "jpeg"
	case "image/png":
		return "png"
	case "image/webp":
		return "webp"
	case "image/gif":
		return "gif"
	case "image/bmp":
		return "bmp"
	case "image/tiff":
		return "tiff"
	case "application/pdf":
		return "pdf"

	default:
		return "unknown"
	}

}
