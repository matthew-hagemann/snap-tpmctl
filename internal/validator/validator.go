package validator

import (
	"fmt"
)

// TODO: use validator.Check ?

func ValidateKey(key int) (int, error) {
	if key < 0 {
		return key, fmt.Errorf("missing key-id argument")
	}

	return key, nil
}
