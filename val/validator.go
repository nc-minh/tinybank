package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidFullname = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("length must be between %d and %d", minLength, maxLength)
	}

	return nil
}

func ValidateUsername(username string) error {
	err := ValidateString(username, 3, 100)
	if err != nil {
		return fmt.Errorf("invalid username: %w", err)
	}

	if !isValidUsername(username) {
		return fmt.Errorf("username must contain only lowercase letters, digits and underscores")
	}

	return nil
}

func ValidateFullname(username string) error {
	err := ValidateString(username, 3, 100)
	if err != nil {
		return fmt.Errorf("invalid username: %w", err)
	}

	if !isValidFullname(username) {
		return fmt.Errorf("fullname must contain only lowercase letters or spaces")
	}

	return nil
}

func ValidatePassword(password string) error {
	return ValidateString(password, 6, 100)
}

func ValidateEmail(email string) error {
	if err := ValidateString(email, 6, 100); err != nil {
		return fmt.Errorf("invalid email: %w", err)
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return fmt.Errorf("invalid email: %w", err)
	}

	return nil
}

func ValidateEmailId(value int64) error {
	if value <= 0 {
		return fmt.Errorf("invalid email id")
	}

	return nil
}

func ValidateSecretCode(value string) error {
	return ValidateString(value, 32, 128)
}
