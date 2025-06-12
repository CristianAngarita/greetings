package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Nombre vacío")
	}
	msg := fmt.Sprintf(randomFormat(), name)
	return msg, nil
}

func randomFormat() string {

	formats := []string{
		"Hola %v Bienvenido",
		"Que bueno verte %v",
		"Saludo %v Encantado",
	}
	return formats[rand.Intn(len(formats))]
}

func Hellos(names []string) (map[string]string, error) {
	msgs := make(map[string]string)
	for _, name := range names {
		msg, err := Hello(name)
		if err != nil {
			return nil, err
		}
		msgs[name] = msg
	}
	return msgs, nil
}
