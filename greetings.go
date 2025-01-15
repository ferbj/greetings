package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// hello
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Name must not be empty")
	}
	message := fmt.Sprintf(randomGreetings(), name)
	return message, nil
}

func Grettings(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomGreetings() string {
	greetings := []string{
		"Hola, %v Welcome!",
		"Welcome, %v nice to meet you, have a nice day",
		"Hi, %v wish you a good day",
		"Saludos, %v te deseo buen dia",
		"Hallo, %v ich hoffe bist du gut",
		"Bonjour, %v es-tu bien?",
	}
	//return random greetings
	return greetings[rand.Intn(len(greetings))]
}
