package main

var justString string

func someFunc() {
	v := createHugeString(1 << 10) // создаем строку размером 1024 символа
	justString = v[:100]           // берем только первые 100 символов
}

// т.к строка под капотом слайс байт, то в ней хранится указатель на родительский массив размером 1024 байта
// justString - срез из 100 байт, содержит ссылку на массив из 1024 байта
// пока существует justString, массив из 1024 байта не будет удален сборщиком мусора, получается утечка памяти
// решение - копировать нужные данные в новый слайс

func improvedSomeFunc() {
	v := createHugeString(1 << 10)   // создаем строку размером 1024 символа
	tmp := v[:100]                   // берем только первые 100 символов
	justString = string([]byte(tmp)) // копируем данные в новый слайс
	// можно еще через string.Clone strings.Clone(tmp)
}
func main() {
	improvedSomeFunc()
}

func createHugeString(size int) string {
	return string(make([]byte, size))
}
