package newpackage

import (
	"gopackages/wordz"
)

func City() string {
	wordz.Prefix = "Random city: "
	wordz.Words = []string{"Анадырь", "Багратионовск", "Верхний Уфалей", "Горнозаводск", "Дно", "Емва", "Зея"}

	return wordz.Random()

}
