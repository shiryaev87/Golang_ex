package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// Реализуем метод Read для нашей структуры
func (rot rot13Reader) Read(b []byte) (int, error) {
	// 1. Просим внутренний источник (r) налить "воды" (байт) в наше ведро (b)
	n, err := rot.r.Read(b)

	// 2. Проходимся только по тем байтам, которые удалось прочитать
	for i := 0; i < n; i++ {
		char := b[i]

		// 3. Применяем шифр ROT13: сдвигаем букву на 13 позиций по алфавиту
		if (char >= 'A' && char <= 'M') || (char >= 'a' && char <= 'm') {
			b[i] += 13
		} else if (char >= 'N' && char <= 'Z') || (char >= 'n' && char <= 'z') {
			b[i] -= 13
		}
		// Все остальные символы (пробелы, знаки препинания) оставляем как есть
	}

	// 4. Возвращаем количество прочитанных байт и статус (например, io.EOF)
	return n, err
}

func main() {
	// Скрытое сообщение: "You cracked the code!"
	s := strings.NewReader("Lbh penpxrq gur pbqr!")

	// Оборачиваем строку в наш "фильтр"
	r := rot13Reader{s}

	// io.Copy будет выкачивать данные из r, пока не встретит ошибку io.EOF
	io.Copy(os.Stdout, &r)
}
