package game_application

import (
	"errors"
	"strconv"
)

func InsertDelimeter(s string) string {
	var buff string
	for _, c := range s {
		buff += string(c)
		buff += string(',')
	}
	return buff
}

func IsOperator(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/"
}

func GetSliceOfOperators(ss []string) []int {
	var result []int
	for _, s := range ss {
		_, sign := strconv.Atoi(s)
		if sign != nil {
			switch {
			case s[0] == '+':
				result = append(result, '+')
			case s[0] == '-':
				result = append(result, '-')
			case s[0] == '*':
				result = append(result, '*')
			case s[0] == '/':
				result = append(result, '/')
			}
		} else {
			result = append(result, -1)
		}
	}
	return result
}

func GetSliceOfNumbers(ss []string) []int {
	var result []int
	for _, s := range ss {
		v, op := strconv.Atoi(s)
		if op != nil {
			result = append(result, -1)
		} else {
			result = append(result, v)
		}
	}
	return result
}

func InitGameStructure(ss []string) ([]int, []int, error) {
	for _, s := range ss {
		if !IsValidArg(int(s[0])) {
			return nil, nil, errors.New("invalid argument")
		}
	}
	return GetSliceOfNumbers(ss), GetSliceOfOperators(ss), nil
}

func IsAvaiableInSolution(b rune, solution string) bool {
	for _, c := range solution {
		if b == c {
			return true
		}
	}
	return false
}

func IsInRightSpot(try rune, solution rune) bool {
	return try == solution
}

func GetHints(try string, solution string) string {
	var hints [6]byte
	for i, c := range try {
		if IsAvaiableInSolution(c, solution) {
			if IsInRightSpot(c, rune(solution[i])) {
				hints[i] = 'C'
			} else {
				hints[i] = 'T'
			}
		} else {
			hints[i] = 'X'
		}
	}
	return string(hints[:])
}
