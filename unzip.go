package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// Определяем пути
	inputDir := "INPUT_FOLDER"
	outputDir := "."

	// Получаем список файлов в папке input
	files, err := os.ReadDir(inputDir)
	if err != nil {
		fmt.Printf("Ошибка чтения папки %s: %v\n", inputDir, err)
		return
	}

	// Обрабатываем каждый файл
	for _, file := range files {
		if file.IsDir() {
			continue // Пропускаем подпапки
		}

		// Проверяем расширение .zip
		if filepath.Ext(file.Name()) == ".zip" {
			zipPath := filepath.Join(inputDir, file.Name())
			fmt.Printf("Распаковка %s...\n", zipPath)

			// Распаковываем архив
			err := unzip(zipPath, outputDir)
			if err != nil {
				fmt.Printf("Ошибка распаковки %s: %v\n", zipPath, err)
			} else {
				fmt.Printf("Успешно распакован: %s\n", zipPath)
			}
		}
	}

	fmt.Println("Готово!")
}

// Функция для распаковки ZIP-архива
func unzip(src, dest string) error {
	// Открываем архив
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	// Создаем папку назначения, если её нет
	if err := os.MkdirAll(dest, 0755); err != nil {
		return err
	}

	// Обходим все файлы в архиве
	for _, f := range r.File {
		// Создаем полный путь для файла
		fpath := filepath.Join(dest, f.Name)

		// Проверяем, не является ли файл директорией
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, 0755)
			continue
		}

		// Создаем родительские папки, если нужно
		if err := os.MkdirAll(filepath.Dir(fpath), 0755); err != nil {
			return err
		}

		// Открываем файл в архиве
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		// Создаем файл на диске
		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer outFile.Close()

		// Копируем содержимое
		_, err = io.Copy(outFile, rc)
		if err != nil {
			return err
		}
	}

	return nil
}
