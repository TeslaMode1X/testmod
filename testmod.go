package testmod

import (
	"errors"
	"fmt"
)

func Hi(name, lang string) (string, error) {
	if name == "" {
		return "", errors.New("name cannot be empty")
	}

	switch lang {
	case "en":
		return fmt.Sprintf("Hello, %s!", name), nil
	case "pt":
		return fmt.Sprintf("Olá, %s!", name), nil
	default:
		return "", errors.New("unsupported language")
	}
}
