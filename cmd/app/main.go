package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

func main() {
	words := map[string]string{
		"яблоко":    "apple",
		"апельсин":  "orange",
		"банан":     "banana",
		"вишня":     "cherry",
		"киви":      "kiwi",
		"персик":    "peach",
		"ананас":    "pineapple",
		"малина":    "raspberry",
		"грейпфрут": "grapefruit",
		"слива":     "plum",
	}

	random_words := []string{}
	for key := range words {
		random_words = append(random_words, key)
	}

	rand.Shuffle(len(random_words), func(i, j int) {
		random_words[i], random_words[j] = random_words[j], random_words[i]
	})

	score := 0
	number := 1
	Errors := map[string]string{}

	fmt.Fprintln(os.Stdout)
	for _, russian_word := range random_words {
		attempts := 1
		fmt.Fprintf(os.Stdout, "\tСлово № %d\n", number)
		fmt.Fprintf(os.Stdout, "Переведи слово: %s\n", russian_word)
		number++

		for {
			fmt.Fprint(os.Stdout, "Ответ: ")
			ui := bufio.NewReader(os.Stdin)
			translation, err := ui.ReadString('\n')
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка ввода\nПопробуйте еще раз")
				fmt.Fprintln(os.Stdout)
				continue
			}

			translation = strings.TrimSpace(translation)

			if strings.EqualFold(translation, words[russian_word]) {
				fmt.Fprint(os.Stdout, "Верно!\n\n")
				score++
				break
			} else {
				if attempts == 0 {
					fmt.Fprintf(os.Stdout, "Правильный ответ: %s\n\n", words[russian_word])
					Errors[russian_word] = words[russian_word]
					break
				} else {
					fmt.Fprint(os.Stdout, "Попробуй еще раз.\n")
				}

				attempts--
			}
		}
	}

	fmt.Fprintf(os.Stdout, "Правильных ответов: %d из %d\n\n", score, len(words))

	if len(Errors) != 0 {
		fmt.Fprint(os.Stdout, "Непереведенные слова:\n\n")
		for key, value := range Errors {
			fmt.Fprintf(os.Stdout, "%s - %s\n", key, value)
		}
	} else {
		fmt.Fprint(os.Stdout, "Ошибок нет!\n")
	}
	fmt.Fprintln(os.Stdout)
}
