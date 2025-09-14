package converter_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	converter "github.com/v-grabko1999/converter"
	_ "github.com/v-grabko1999/converter/drivers/bmp"
	_ "github.com/v-grabko1999/converter/drivers/jpg"
	_ "github.com/v-grabko1999/converter/drivers/pdf"
	_ "github.com/v-grabko1999/converter/drivers/webp"
)

func TestConvertFilesToPNG(t *testing.T) {
	// Шлях до файлів для тестування
	files := []string{
		"./test/1.bmp",
		"./test/1.jpeg",
		"./test/1.jpg",
		"./test/1.pdf",
		"./test/1.webp",
	}

	// Перебираємо кожен файл для тестування
	for _, filePath := range files {
		// Читаємо байтовий вміст файлу
		inputBytes, err := os.ReadFile(filePath)
		if err != nil {
			t.Fatalf("Помилка при читанні файлу: %v", err)
		}

		// Виконуємо конвертацію файлу у PNG за допомогою функції Convert
		outputBytes, err := converter.Convert(inputBytes, converter.DetectFormat(inputBytes))
		if err != nil {
			t.Fatalf("%s-> Помилка при конвертації у PNG: %v", filePath, err)
		}

		// Генеруємо шлях для збереження результату
		outputPath := fmt.Sprintf("%s_converted.png", filePath)
		err = os.WriteFile(outputPath, outputBytes, 0644)
		if err != nil {
			t.Fatalf("Помилка при записі PNG файлу: %v", err)
		}

		defer os.Remove(outputPath)

		fmt.Printf("Конвертація успішна. PNG файл збережено за шляхом: %s\n", outputPath)

		// Додатково перевіряємо, що результат не порожній
		assert.NotEmpty(t, outputBytes, "PNG файл не повинен бути порожнім")
	}
}
