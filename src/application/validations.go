package game_application

import (
	"errors"
	"fmt"
	"unicode"
)

func ValidateNumArgs(av []string) error {
	if len(av) != 6 {
		return errors.New("invalid number of arguments")
	}
	return nil
}

func CheckArgLen(s string) error {
	if len(s) > 1 {
		return errors.New("argument have more than two digits")
	}
	return nil
}

func IsDigit(ch int) bool {
	return unicode.IsDigit(rune(ch))
}
func IsValidArg(ch int) bool {
	return (ch == '+' || ch == '-' || ch == '*' || ch == '/') || IsDigit(ch)
}

func CheckFirstSpot(ss []string) error {
	first := ss[0][0]
	switch first {
	case '*', '/':
		return errors.New("operator (* and /) shouldt be in the first spot")
	}
	return nil
}

func CheckLastSpot(ss []string) error {
	last := ss[len(ss)-1][0]
	switch last {
	case '+', '-', '*', '/':
		return errors.New("should not use any operator in the last spot")
	}
	return nil
}

func ValidateEntry(av []string) error {
	for _, s := range av {
		if s == "" {
			return errors.New("invalid number of arguments")
		}
	}
	if err := ValidateNumArgs(av); err != nil {
		return err
	}
	if err := CheckFirstSpot(av); err != nil {
		return err
	}
	if err := CheckLastSpot(av); err != nil {
		return err
	}
	return nil
}

func CheckInvalidSequence(av []string) error {
	i := 1
	for i < len(av)-1 {
		if (av[i][0] == '/' && av[i-1][0] == '+') || (av[i][0] == '/' && av[i-1][0] == '-') {
			return fmt.Errorf("%c cant be followed by %c", av[i-1][0], av[i][0])
		}
		if (av[i][0] == '*' && av[i-1][0] == '+') || (av[i][0] == '*' && av[i-1][0] == '-') {
			return fmt.Errorf("%c cant be followed by %c", av[i-1][0], av[i][0])
		}
		if (av[i][0] == '*' && av[i+1][0] == '/') || (av[i][0] == '/' && av[i+1][0] == '*') {
			return fmt.Errorf("%c cant be followed by %c", av[i][0], av[i+1][0])
		}
		i++
	}
	return nil
}

func CheckOperators(av []string) error {
	if err := CheckInvalidSequence(av); err != nil {
		return err
	}
	temp := av[0][0]
	i := 1
	for i < len(av) {
		if IsOperator(av[i]) {
			if av[i][0] == temp {
				return fmt.Errorf("%c cant be followed by %c", temp, av[i][0])
			}
		}
		temp = av[i][0]
		i++
	}
	return nil
}

func ValidateArguments(av []string) error {
	if err := ValidateEntry(av); err != nil {
		return err
	}
	if err := CheckOperators(av); err != nil {
		return err
	}
	for _, s := range av {
		if err := CheckArgLen(s); err != nil {
			return err
		}
		for _, ch := range s {
			if !IsValidArg(int(ch)) {
				return fmt.Errorf("argument `%c` is invalid", ch)
			}
		}
	}
	return nil
}
