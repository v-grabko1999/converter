//go:build windows
// +build windows

package pdf

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os"
	"os/exec"
	"path/filepath"
)

// checkPoppler перевіряє наявність бібліотеки Poppler для Windows
func checkLibs() {
	_, err := exec.LookPath("pdftoppm")
	if err != nil {
		panic(fmt.Sprintf("Poppler is not installed: %v. Please install it to proceed.", err))
	}
}

// convert - функція для Windows, яка конвертує PDF у PNG за допомогою Poppler
func (d *PDFDriver) convert(file []byte) (image.Image, error) {
	// Генеруємо унікальне ім'я файлу для фізичного диска
	tempFileName := generateUniqueFileName("pdf")
	tempDir := os.TempDir()
	tempFilePath := filepath.Join(tempDir, tempFileName)

	err := os.WriteFile(tempFilePath, file, 0644)
	if err != nil {
		return nil, fmt.Errorf("error writing file to disk: %v", err)
	}
	defer os.Remove(tempFilePath) // Видаляємо файл після використання

	// Викликаємо Poppler для конвертації PDF у PNG
	outputFileBase := generateUniqueFileName("png") // Генеруємо унікальне ім'я для PNG файлу
	outputFile := filepath.Join(tempDir, outputFileBase)
	cmd := exec.Command("pdftoppm", "-png", tempFilePath, outputFile)
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("error running pdftoppm: %v", err)
	}

	// Читаємо згенероване PNG з першої сторінки
	pngFilePath := fmt.Sprintf("%s-1.png", outputFile)
	pngFileBytes, err := os.ReadFile(pngFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading PNG file: %v", err)
	}
	defer os.Remove(pngFilePath) // Видаляємо PNG після використання

	// Декодуємо PNG у зображення
	img, err := png.Decode(bytes.NewReader(pngFileBytes))
	if err != nil {
		return nil, fmt.Errorf("error decoding PNG file: %v", err)
	}

	return img, nil
}
