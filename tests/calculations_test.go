package main

import (
	"strings"
	"testing"

	game_application "database_golang/src/application"
)

func TestCheckNumberOfOperations_expect_given_number_of_operations_for_each_test(t *testing.T) {
	expected := []int{0, 1, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2}
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 1 5 5 2 4", " ")},
		{strings.Split("0 1 + 5 2 4", " ")},
		{strings.Split("+ 0 0 0 0 0", " ")},
		{strings.Split("- 0 0 0 0 1", " ")},
		{strings.Split("0 0 1 0 - 1", " ")},
		{strings.Split("1 + - 1 0 0", " ")},
		{strings.Split("1 - + 1 0 0", " ")},
		{strings.Split("- + 1 0 0 0", " ")},
		{strings.Split("+ - 1 0 0 0", " ")},
		{strings.Split("- 1 0 + 1 0", " ")},
		{strings.Split("- 1 0 + - 5", " ")},
		{strings.Split("- 1 + - 5 5", " ")},
		{strings.Split("- 1 * - 5 5", " ")},
		{strings.Split("- 1 / - 5 5", " ")},
		{strings.Split("1 0 0 * 0 5", " ")},
		{strings.Split("1 0 0 0 / 5", " ")},
		{strings.Split("- 1 0 0 / 5", " ")},
		{strings.Split("- 1 0 / - 5", " ")},
		{strings.Split("- 1 + 1 + 5", " ")},
		{strings.Split("+ 1 + 1 + 5", " ")},
		{strings.Split("1 + 1 + 4 5", " ")},
		{strings.Split("1 + - 1 + 1", " ")},
		{strings.Split("1 * - 1 - 1", " ")},
	}
	for i, tt := range tests {
		_, operators, _ := game_application.InitGameStructure(tt.args)
		got := game_application.CheckNumberOfOperations(operators)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}

func TestGetValue_passing_equation_expect_int_representation_of_that_number(t *testing.T) {
	expected := []int{115524, 15524, 0, 1, -1, 0, 0, 1, 1, 10, 100, 1000, -100, -10, -1, 100, 10, 1, -1000, -1000, 1, -1, 10, -10, 100, -100, -1, -1}
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 1 5 5 2 4", " ")},
		{strings.Split("0 1 5 5 2 4", " ")},
		{strings.Split("0 0 0 0 0 0", " ")},
		{strings.Split("0 0 0 0 0 1", " ")},
		{strings.Split("- 0 0 0 0 1", " ")},
		{strings.Split("- 0 0 0 0 0", " ")},
		{strings.Split("+ 0 0 0 0 0", " ")},
		{strings.Split("+ 0 0 0 0 1", " ")},
		{strings.Split("1 + 0 0 0 0", " ")},
		{strings.Split("1 0 - 0 0 0", " ")},
		{strings.Split("1 0 0 * 0 0", " ")},
		{strings.Split("1 0 0 0 / 0", " ")},
		{strings.Split("- 1 0 0 / 0", " ")},
		{strings.Split("- 1 0 * 0 0", " ")},
		{strings.Split("- 1 + 0 0 0", " ")},
		{strings.Split("+ 1 0 0 / 0", " ")},
		{strings.Split("+ 1 0 * 0 0", " ")},
		{strings.Split("+ 1 + 0 0 0", " ")},
		{strings.Split("- + 1 0 0 0", " ")},
		{strings.Split("+ - 1 0 0 0", " ")},
		{strings.Split("+ 1 + 0 0 0", " ")},
		{strings.Split("- 1 + 0 0 0", " ")},
		{strings.Split("+ 1 0 * 0 0", " ")},
		{strings.Split("- 1 0 * 0 0", " ")},
		{strings.Split("+ 1 0 0 / 0", " ")},
		{strings.Split("- 1 0 0 / 0", " ")},
		{strings.Split("- + 0 1 / 0", " ")},
		{strings.Split("+ - 0 1 / 0", " ")},
	}
	for i, tt := range tests {
		numbers, operators, _ := game_application.InitGameStructure(tt.args)
		got := game_application.GetValue(numbers, operators)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}

func TestCalculateSingle_passing_single_operations_expect_correct_results(t *testing.T) {
	expected := []int{5525, 1555, 1555, 40, 2000, 0, -25, 5, 0, 25, 50, -1, -2, -20, 2, -5, -15, 95, 5, -1, 5, -15, -15, -10, -10}
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 + 5 5 2 4", " ")},
		{strings.Split("1 5 5 5 + 0", " ")},
		{strings.Split("1 5 5 5 - 0", " ")},
		{strings.Split("1 0 * 0 0 4", " ")},
		{strings.Split("1 0 * 2 0 0", " ")},
		{strings.Split("1 0 * 0 0 0", " ")},
		{strings.Split("- 5 * 0 0 5", " ")},
		{strings.Split("2 5 / 0 0 5", " ")},
		{strings.Split("0 1 / 1 0 0", " ")},
		{strings.Split("- 5 * - 0 5", " ")},
		{strings.Split("- 1 0 * - 5", " ")},
		{strings.Split("- 0 0 1 * 1", " ")},
		{strings.Split("- 0 1 0 / 5", " ")},
		{strings.Split("1 0 0 / - 5", " ")},
		{strings.Split("- 1 0 / - 5", " ")},
		{strings.Split("- 1 0 + 0 5", " ")},
		{strings.Split("- 1 0 + - 5", " ")},
		{strings.Split("1 0 0 + - 5", " ")},
		{strings.Split("+ 0 0 + 0 5", " ")},
		{strings.Split("+ - 6 + 0 5", " ")},
		{strings.Split("+ 1 0 + - 5", " ")},
		{strings.Split("+ - 1 0 - 5", " ")},
		{strings.Split("- + 1 0 - 5", " ")},
		{strings.Split("- + 5 + - 5", " ")},
		{strings.Split("+ - 5 + - 5", " ")},
	}
	for i, tt := range tests {
		numbers, operators, _ := game_application.InitGameStructure(tt.args)
		got, _ := game_application.CalculateSingle(numbers, operators)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}

func TestCalculateSingle_DivisionByZero_expect_error(t *testing.T) {
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 / 0 0 0 0", " ")},
		{strings.Split("1 0 / 0 0 0", " ")},
		{strings.Split("1 0 0 / 0 0", " ")},
		{strings.Split("1 0 0 0 / 0", " ")},
		{strings.Split("0 0 / 0 0 0", " ")},
		{strings.Split("- 1 / 0 0 0", " ")},
		{strings.Split("- 1 0 / 0 0", " ")},
		{strings.Split("- 1 0 0 / 0", " ")},
		{strings.Split("- 1 0 / - 0", " ")},
		{strings.Split("1 0 0 / - 0", " ")},
		{strings.Split("1 / - 0 0 0", " ")},
		{strings.Split("- + 1 / 0 0", " ")},
		{strings.Split("+ - 1 / 0 0", " ")},
	}
	for i, tt := range tests {
		numbers, operators, _ := game_application.InitGameStructure(tt.args)
		_, err := game_application.CalculateSingle(numbers, operators)
		if err == nil {
			t.Errorf("test %d: expected error. didnt get one", i)
		}
	}
}

func TestIsPrecedence_passing_equation_with_and_without_precedence_expect_true_or_false(t *testing.T) {
	expected := []bool{true, false, false, true, true, false, false, false, true}
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 1 + 5 * 4", " ")},
		{strings.Split("1 + 2 + - 3", " ")},
		{strings.Split("- 1 + 5 - 2", " ")},
		{strings.Split("- 1 + 5 / 2", " ")},
		{strings.Split("1 + 1 * 1 2", " ")},
		{strings.Split("1 + 1 + 4 0", " ")},
		{strings.Split("2 / 1 + 4 0", " ")},
		{strings.Split("2 / - 1 + 2", " ")},
		{strings.Split("2 + 2 / - 2", " ")},
	}
	for i, tt := range tests {
		_, operators, _ := game_application.InitGameStructure(tt.args)
		got := game_application.IsPrecedence(operators)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}
func TestGetPrecedenceIndex_passing_precedence_operator_expect_its_index(t *testing.T) {
	expected := []int{4, 3, 3, 3, 4}
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 1 + 5 * 4", " ")},
		{strings.Split("1 1 2 / 2 4", " ")},
		{strings.Split("+ 1 2 / - 3", " ")},
		{strings.Split("1 + 2 * - 3", " ")},
		{strings.Split("- 2 - 1 / 3", " ")},
	}
	for i, tt := range tests {
		_, operators, _ := game_application.InitGameStructure(tt.args)
		got := game_application.GetPrecedenceIndex(operators)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}

func TestCalculate_passing_valid_equations_expect_its_results(t *testing.T) {
	expected := []int{0, -3, 4, 17, 0, 0, -6, -5, 10, -8, 2, 11, 1, 1, -3, 20, 40, -9, -9, 0}
	var tests = []struct {
		args []string
	}{
		{strings.Split("- 1 + 1 + 0", " ")},
		{strings.Split("- 1 * 2 - 1", " ")},
		{strings.Split("- 1 + 2 + 3", " ")},
		{strings.Split("1 0 + 5 + 2", " ")},
		{strings.Split("+ 6 / 2 - 3", " ")},
		{strings.Split("1 + 2 + - 3", " ")},
		{strings.Split("- 2 - 1 - 3", " ")},
		{strings.Split("- 2 * 1 - 3", " ")},
		{strings.Split("6 / 3 + 0 8", " ")},
		{strings.Split("- 2 * 1 * 4", " ")},
		{strings.Split("0 + 1 + 0 1", " ")},
		{strings.Split("1 0 + 0 + 1", " ")},
		{strings.Split("+ 1 + 0 + 0", " ")},
		{strings.Split("+ 0 + 0 + 1", " ")},
		{strings.Split("- 1 - 1 - 1", " ")},
		{strings.Split("1 * 2 0 - 0", " ")},
		{strings.Split("1 * 2 0 * 2", " ")},
		{strings.Split("- 6 / 2 * 3", " ")},
		{strings.Split("- 6 * 3 / 2", " ")},
		{strings.Split("+ 6 / 2 - 3", " ")},
	}
	for i, tt := range tests {
		numbers, operators, _ := game_application.InitGameStructure(tt.args)
		got, _ := game_application.Calculate(numbers, operators)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}

func TestCalculate_with_precedence_passing_valid_equations_expect_its_results(t *testing.T) {
	expected := []int{45, 1, 3, 105, 109, 3, -1, -3, 42}
	var tests = []struct {
		args []string
	}{
		{strings.Split("5 + 2 * 2 0", " ")},
		{strings.Split("6 - 5 / 0 1", " ")},
		{strings.Split("+ 1 5 / 0 5", " ")},
		{strings.Split("5 + 5 0 * 2", " ")},
		{strings.Split("9 + 5 0 * 2", " ")},
		{strings.Split("- 1 + 2 * 2", " ")},
		{strings.Split("- 2 + 2 / 2", " ")},
		{strings.Split("+ 2 - 1 * 5", " ")},
		{strings.Split("2 + 2 * 2 0", " ")},
	}
	for i, tt := range tests {
		numbers, operators, _ := game_application.InitGameStructure(tt.args)
		got, _ := game_application.Calculate(numbers, operators)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}

func TestCalculate_with_precedence_passing_invalid_equations_expect_DivisionByZero_error(t *testing.T) {
	var tests = []struct {
		args []string
	}{
		{strings.Split("5 + 2 / 0 0", " ")},
		{strings.Split("6 - 5 / 0 0", " ")},
		{strings.Split("- 1 * 2 / 0", " ")},
		{strings.Split("+ 1 + 1 / 0", " ")},
		{strings.Split("5 + 5 0 / 0", " ")},
		{strings.Split("- 1 + 2 / 0", " ")},
		{strings.Split("- 2 / 2 / 0", " ")},
		{strings.Split("+ 2 - 1 / 0", " ")},
		{strings.Split("2 * 2 / 0 0", " ")},
		{strings.Split("2 / 2 / 0 0", " ")},
	}
	for i, tt := range tests {
		numbers, operators, _ := game_application.InitGameStructure(tt.args)
		_, err := game_application.Calculate(numbers, operators)
		if err == nil {
			t.Errorf("test %d: expect an error but we didnt get one", i)
		}
	}
}
