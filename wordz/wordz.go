package wordz

import "math/rand"
import "time"
import "fmt" //Не забываем импорт пакета fmt

var Hello = "This is wordz package"
var Prefix = "Random word: "
var Words = []string{"One", "Two", "Three", "Four", "Five"}

func init() { // Добавили функцию. Она сработает при импорте пакета wordz в файле main.go
	fmt.Println("Function init in package wordz")
}

func Random() string {
	max := len(Words)
	rand.Seed(time.Now().Unix())
	return get(rand.Intn(max))
}

func get(index int) string {
	return Prefix + Words[index]
}


