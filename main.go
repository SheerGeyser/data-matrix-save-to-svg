package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/boombuler/barcode/datamatrix"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите текст (или отсканируйте код). Для выхода введите 'exit':")

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		text := strings.TrimSpace(input)
		if strings.ToLower(text) == "exit" {
			break
		}

		err = generateSVG(text)
		if err != nil {
			log.Println("Ошибка генерации:", err)
		} else {
			fmt.Println("SVG сохранён.")
		}
	}
}

func generateSVG(data string) error {
	code, err := datamatrix.Encode(data)
	if err != nil {
		return err
	}

	bounds := code.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	filename := fmt.Sprintf("%s.svg", sanitizeFilename(data))
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	cellSize := 10

	svgWidth := width * cellSize
	svgHeight := height * cellSize

	fmt.Fprintf(file, `<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d">`, svgWidth, svgHeight)
	fmt.Fprintln(file)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			gray := code.At(x, y)
			r, g, b, _ := gray.RGBA()
			if r == 0 && g == 0 && b == 0 {
				fmt.Fprintf(file,
					`  <rect x="%d" y="%d" width="%d" height="%d" fill="black" />`+"\n",
					x*cellSize, y*cellSize, cellSize, cellSize)
			}
		}
	}

	fmt.Fprintln(file, "</svg>")

	return nil
}

func sanitizeFilename(name string) string {
	return strings.Map(func(r rune) rune {
		if r == '_' || r == '-' || r == '.' ||
			(r >= 'a' && r <= 'z') ||
			(r >= 'A' && r <= 'Z') ||
			(r >= '0' && r <= '9') {
			return r
		}
		return '_'
	}, name)
}
