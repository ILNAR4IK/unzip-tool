# ZIP Archive Extractor

Простая утилита на Go для распаковки всех ZIP-архивов из папки `input` в корневую директорию, с автоматическим удалением избыточной вложенности папок.

## Особенности

- Распаковывает все ZIP-архивы из папки `INPUT_FOLDER`
- Автоматически пропускает первый уровень вложенности (если архив содержит одну основную папку)
- Сохраняет структуру файлов внутри архива
- Логирование процесса в консоль

## Требования

- Установленный [Go](https://golang.org/dl/) (версия 1.16+)

## Установка

1. Склонируйте репозиторий:
```bash
git clone https://github.com/ILNAR4IK/unzip-tool.git
cd unzip-tool