package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}

	//var message string
	//var message = fmt.Sprintf(randomFormat(), name)
	var message = fmt.Sprintf(randomFormat())
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		var message, err = Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	//var formats = []string
	formats := []string{
		"(1)Hi",
		"(2)Swasdee",
		"(3)Good Morning",
	}

	/* formats := []string{
		"(1)Hi, %v. Welcome!",
		"(2)Great to see you, %v!",
		"(3)Hail, %v! Well met!",
	} */

	return formats[rand.Intn(len(formats))]
}
