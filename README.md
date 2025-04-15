# DataMatrix SVG Generator (CLI)

Простая терминальная программа на Go, которая:

- Принимает ввод с клавиатуры (или сканера Data Matrix).
- Генерирует и сохраняет SVG-файл с DataMatrix-кодом.
- Подходит для автоматизации, терминалов, встроенных систем.

---

## 🔧 Установка

### Клонируйте репозиторий:

```bash
git clone git@github.com:your-username/datamatrix-svg.git
cd datamatrix-svg
```

### Сборка:

#### Под текущую ОС:

```bash
go build -o datamatrix
```

#### Или для других платформ (пример для Linux и macOS):

```bash
GOOS=linux GOARCH=amd64 go build -o datamatrix-linux
GOOS=darwin GOARCH=arm64 go build -o datamatrix-macos-arm64
```

---

## 🚀 Использование

Запуск из терминала или [скачать исполняемый файл в последнем релизе](https://github.com/SheerGeyser/data-matrix-save-to-svg/releases/tag/v1.0.0):

```bash
./datamatrix
```

После запуска программа будет ожидать ввод:

```
Введите текст (или отсканируйте код). Для выхода введите 'exit':
> HelloWorld
SVG сохранён.
> exit
```

На каждый введённый текст будет сгенерирован SVG-файл рядом с программой, например:

```
HelloWorld.svg
```

---

## 📦 Пример использования

Можно подключить сканер, настроенный на ввод текста + `Enter`. Сразу после сканирования программа сгенерирует SVG.

Это удобно, например:

- В логистике (маркировка коробок)
- Для внутреннего учёта
- При подготовке печати QR/DataMatrix кодов

---

## 🛠 Зависимости

Используется библиотека:

- [`github.com/boombuler/barcode`](https://github.com/boombuler/barcode) — генерация Data Matrix.

---

## 🧹 Возможности доработки

- Поддержка разных типов штрихкодов (QR, Code128 и т.п.).
- Настройки размеров и цветов SVG.
- Генерация PNG/JPG в дополнение к SVG.
- GUI-обёртка (через WebView или Wails).

---
