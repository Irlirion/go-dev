package unicode_name

import "unicode"

// AnalyzeName считает байты, Unicode-кодовые точки, латинские и кириллические
// буквы и возвращает байтовые позиции начала каждой Unicode-кодовой точки.
// Предполагается, что входная строка содержит корректный UTF-8.
func AnalyzeName(name string) (byteCount, runeCount, latinCount, cyrillicCount int, positions []int) {
	byteCount = len(name)
	for i, letter := range name {
		runeCount++
		positions = append(positions, i)
		if unicode.IsLetter(letter) && unicode.Is(unicode.Latin, letter) {
			latinCount++
		}
		if unicode.IsLetter(letter) && unicode.Is(unicode.Cyrillic, letter) {
			cyrillicCount++
		}
	}
	return byteCount, runeCount, latinCount, cyrillicCount, positions
}
