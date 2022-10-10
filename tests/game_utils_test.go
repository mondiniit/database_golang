package main

import (
	"strings"
	"testing"

	game_application "database_golang/src/application"
)

func TestIsOperator_passing_valid_operator_expect_true(t *testing.T) {
	operators := []string{"-", "+", "/", "*"}
	for i, s := range operators {
		op := game_application.IsOperator(s)
		if op != true {
			t.Errorf("test %d: expect true but got false", i)
		}
	}
}

func TestIsOperator_passing_invalid_operator_expect_false(t *testing.T) {
	operators := []string{">", "<", "=", ".", "x", "(", ")", "{", "}", "1", "2", "9", "0"}
	for i, s := range operators {
		op := game_application.IsOperator(s)
		if op != false {
			t.Errorf("test %d: expect false but got true", i)
		}
	}
}

func TestGetSliceOfOperators_passing_numbers_only_expect_slice_of_minus_ones(t *testing.T) {
	var results = []struct {
		args []int
	}{
		{[]int{-1, -1, -1, -1, -1, -1}},
		{[]int{-1, -1, -1, -1, -1, -1}},
		{[]int{-1, -1, -1, -1, -1, -1}},
	}
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 1 1 1 1 1", " ")},
		{strings.Split("0 1 2 3 4 5", " ")},
		{strings.Split("9 8 7 6 5 4", " ")},
	}
	for i, tt := range tests {
		operators := game_application.GetSliceOfOperators(tt.args)
		for j, v := range operators {
			if v != results[i].args[j] {
				t.Errorf("test %d: expect %d but got %d", i, results[i].args[j], v)
			}
		}
	}
}

func TestGetSliceOfOperators_passing_operators_and_numbers_expect_ascii_and_minus_ones(t *testing.T) {
	var results = []struct {
		args []int
	}{
		{[]int{-1, '+', '+', -1, -1, -1}},
		{[]int{-1, -1, '*', -1, '-', -1}},
		{[]int{'/', -1, -1, -1, -1, -1}},
		{[]int{'/', -1, -1, -1, -1, '*'}},
		{[]int{-1, -1, -1, -1, -1, '*'}},
		{[]int{'+', -1, '-', '/', '/', -1}},
	}
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 + + 1 1 1", " ")},
		{strings.Split("0 1 * 3 - 5", " ")},
		{strings.Split("/ 8 7 6 5 4", " ")},
		{strings.Split("/ 8 7 6 5 *", " ")},
		{strings.Split("0 8 7 6 5 *", " ")},
		{strings.Split("+ 8 - / / 3", " ")},
	}
	for i, tt := range tests {
		operators := game_application.GetSliceOfOperators(tt.args)
		for j, v := range operators {
			if v != results[i].args[j] {
				t.Errorf("test %d: expect %d but got %d", i, results[i].args[j], v)
			}
		}
	}
}

func TestGetSliceOfOperators_passing_operators_only_expect_its_ascii_values(t *testing.T) {
	var results = []struct {
		args []int
	}{
		{[]int{'+', '-', '*', '/', '*', '+'}},
		{[]int{'-', '+', '+', '+', '/', '+'}},
		{[]int{'-', '-', '-', '-', '-', '-'}},
		{[]int{'-', '/', '*', '-', '-', '-'}},
		{[]int{'/', '+', '+', '+', '/', '/'}},
		{[]int{'-', '*', '/', '+', '*', '*'}},
	}
	var tests = []struct {
		args []string
	}{
		{strings.Split("+ - * / * +", " ")},
		{strings.Split("- + + + / +", " ")},
		{strings.Split("- - - - - -", " ")},
		{strings.Split("- / * - - -", " ")},
		{strings.Split("/ + + + / /", " ")},
		{strings.Split("- * / + * *", " ")},
	}
	for i, tt := range tests {
		operators := game_application.GetSliceOfOperators(tt.args)
		for j, v := range operators {
			if v != results[i].args[j] {
				t.Errorf("test %d: expect %d but got %d", i, results[i].args[j], v)
			}
		}
	}
}

func TestGetSliceOfNumbers_passing_numbers_only_expect_its_integer_representation(t *testing.T) {
	var results = []struct {
		args []int
	}{
		{[]int{1, 0, 1, 0, 4, 0}},
		{[]int{9, 2, 3, 6, 0, 2}},
		{[]int{4, 2, 2, 7, 4, 9}},
	}
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 0 1 0 4 0", " ")},
		{strings.Split("9 2 3 6 0 2", " ")},
		{strings.Split("4 2 2 7 4 9", " ")},
	}
	for i, tt := range tests {
		numbers := game_application.GetSliceOfNumbers(tt.args)
		for j, v := range numbers {
			if v != results[i].args[j] {
				t.Errorf("test %d: expect %d but got %d", i, results[i].args[j], v)
			}
		}
	}
}

func TestGetSliceOfNumbers_passing_operators_only_expect_minus_ones(t *testing.T) {
	var results = []struct {
		args []int
	}{
		{[]int{-1, -1, -1, -1, -1, -1}},
		{[]int{-1, -1, -1, -1, -1, -1}},
		{[]int{-1, -1, -1, -1, -1, -1}},
		{[]int{-1, -1, -1, -1, -1, -1}},
		{[]int{-1, -1, -1, -1, -1, -1}},
	}
	var tests = []struct {
		args []string
	}{
		{strings.Split("+ + + + + +", " ")},
		{strings.Split("- - - - - -", " ")},
		{strings.Split("* * * * * *", " ")},
		{strings.Split("/ / / / / /", " ")},
		{strings.Split("/ * - + / *", " ")},
	}
	for i, tt := range tests {
		numbers := game_application.GetSliceOfNumbers(tt.args)
		for j, v := range numbers {
			if v != results[i].args[j] {
				t.Errorf("test %d: expect %d but got %d", i, results[i].args[j], v)
			}
		}
	}
}
func TestGetSliceOfNumbers_passing_numbers_and_operators_expect_integers_representations_and_minus_ones(t *testing.T) {
	var results = []struct {
		args []int
	}{
		{[]int{1, -1, 1, -1, 4, 0}},
		{[]int{9, -1, 3, -1, 0, 2}},
		{[]int{0, 0, 2, -1, 4, 0}},
	}
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 + 1 + 4 0", " ")},
		{strings.Split("9 - 3 / 0 2", " ")},
		{strings.Split("0 0 2 * 4 0", " ")},
	}
	for i, tt := range tests {
		numbers := game_application.GetSliceOfNumbers(tt.args)
		for j, v := range numbers {
			if v != results[i].args[j] {
				t.Errorf("test %d: expect %d but got %d", i, results[i].args[j], v)
			}
		}
	}
}

func TestInitGameStructure_passing_invalid_args_expect_error(t *testing.T) {
	var tests = []struct {
		args []string
	}{
		{[]string{" ", "0", "0", "0", "0", "0"}},
		{[]string{"0", "0", "0", "0", "0", " "}},
		{[]string{"+", "!", "4", "8", "9", "0"}},
		{[]string{"4", "2", "+", "_", "0", "0"}},
		{[]string{"2", "0", "*", "2", "+", "x"}},
		{[]string{"1", ".", "8", "/", "0", "4"}},
	}
	for i, tt := range tests {
		_, _, err := game_application.InitGameStructure(tt.args)
		if err == nil {
			t.Errorf("test %d: expect error. we didnt get one", i)
		}
	}
}

func TestInitGameStructure_passing_valid_args_expect_no_error(t *testing.T) {
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 + 1 + 4 0", " ")},
		{strings.Split("9 - 3 / 0 2", " ")},
		{strings.Split("0 0 2 * 4 0", " ")},
		{strings.Split("+ + 1 + 4 0", " ")},
		{strings.Split("9 - - - 1 0", " ")},
		{strings.Split("+ - / - * +", " ")},
		{strings.Split("0 0 0 0 0 0", " ")},
		{strings.Split("1 2 4 - 4 2", " ")},
		{strings.Split("2 0 * 2 + 2", " ")},
		{strings.Split("1 6 8 / 0 4", " ")},
		{strings.Split("8 * 8 - 2 2", " ")},
	}
	for i, tt := range tests {
		_, _, err := game_application.InitGameStructure(tt.args)
		if err != nil {
			t.Errorf("test %d: got error `%s` where we didnt expect", i, err)
		}
	}
}

func TestInitGameStructure_passing_operators_only_expect_its_ascii_representations(t *testing.T) {
	var results = []struct {
		args []int
	}{
		{[]int{'+', '-', '*', '/', '*', '+'}},
		{[]int{'-', '+', '+', '+', '/', '+'}},
		{[]int{'-', '-', '-', '-', '-', '-'}},
		{[]int{'-', '/', '*', '-', '-', '-'}},
		{[]int{'/', '+', '+', '+', '/', '/'}},
		{[]int{'-', '*', '/', '+', '*', '*'}},
	}
	var tests = []struct {
		args []string
	}{
		{strings.Split("+ - * / * +", " ")},
		{strings.Split("- + + + / +", " ")},
		{strings.Split("- - - - - -", " ")},
		{strings.Split("- / * - - -", " ")},
		{strings.Split("/ + + + / /", " ")},
		{strings.Split("- * / + * *", " ")},
	}
	for i, tt := range tests {
		_, operators, _ := game_application.InitGameStructure(tt.args)
		for j, v := range operators {
			if v != results[i].args[j] {
				t.Errorf("test %d: expect %d but got %d", i, results[i].args[j], v)
			}
		}
	}
}

func TestInitGameStructure_passing_numbers_only_expect_its_integer_representations(t *testing.T) {
	var results = []struct {
		args []int
	}{
		{[]int{1, 1, 1, 1, 1, 1}},
		{[]int{0, 1, 2, 3, 4, 5}},
		{[]int{9, 8, 7, 6, 5, 4}},
	}
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 1 1 1 1 1", " ")},
		{strings.Split("0 1 2 3 4 5", " ")},
		{strings.Split("9 8 7 6 5 4", " ")},
	}
	for i, tt := range tests {
		numbers, _, _ := game_application.InitGameStructure(tt.args)
		for j, v := range numbers {
			if v != results[i].args[j] {
				t.Errorf("test %d: expect %d but got %d", i, results[i].args[j], v)
			}
		}
	}
}

func TestInitGameStructure_checking_number_slice_expect_its_integer_representations_and_minus_ones_for_op(t *testing.T) {
	var results = []struct {
		args []int
	}{
		{[]int{1, -1, 1, -1, 4, 0}},
		{[]int{9, -1, 3, -1, 0, 2}},
		{[]int{0, 0, 2, -1, 4, 0}},
	}
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 + 1 + 4 0", " ")},
		{strings.Split("9 - 3 / 0 2", " ")},
		{strings.Split("0 0 2 * 4 0", " ")},
	}
	for i, tt := range tests {
		numbers, _, _ := game_application.InitGameStructure(tt.args)
		for j, v := range numbers {
			if v != results[i].args[j] {
				t.Errorf("test %d: expect %d but got %d", i, results[i].args[j], v)
			}
		}
	}
}

func TestInitGameStructure_checking_op_slice_expect_its_ascii_representations_and_minus_ones_for_numbers(t *testing.T) {
	var results = []struct {
		args []int
	}{
		{[]int{-1, '+', -1, '+', -1, -1}},
		{[]int{-1, '-', -1, '/', -1, -1}},
		{[]int{-1, -1, -1, '*', -1, -1}},
	}
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 + 1 + 4 0", " ")},
		{strings.Split("9 - 3 / 0 2", " ")},
		{strings.Split("0 0 2 * 4 0", " ")},
	}
	for i, tt := range tests {
		_, operators, _ := game_application.InitGameStructure(tt.args)
		for j, v := range operators {
			if v != results[i].args[j] {
				t.Errorf("test %d: expect %d but got %d", i, results[i].args[j], v)
			}
		}
	}
}

func TestIsAvaiableInSolution_passing_NA_arg_expect_false(t *testing.T) {
	expected := false
	solution := "022+20"
	var tests = []struct {
		arg rune
	}{
		{'1'}, {'3'}, {'4'}, {'5'}, {'6'}, {'7'},
		{'8'}, {'9'}, {'-'}, {'*'}, {'/'},
	}
	for i, tt := range tests {
		got := game_application.IsAvaiableInSolution(tt.arg, solution)
		if got != expected {
			t.Errorf("test %d: expected %v but got %v", i, expected, got)
		}
	}
}

func TestIsAvaiableInSolution_passing_avaiable_arg_expect_true(t *testing.T) {
	expected := true
	solution := "022+20"
	var tests = []struct {
		arg rune
	}{
		{'0'}, {'2'}, {'+'},
	}
	for i, tt := range tests {
		got := game_application.IsAvaiableInSolution(tt.arg, solution)
		if got != expected {
			t.Errorf("test %d: expected %v but got %v", i, expected, got)
		}
	}
}

func TestIsInRightSpot_passing_arg_in_wrong_spot_expect_false(t *testing.T) {
	expected := false
	solution := "022+20"
	tests := []struct {
		arg rune
	}{
		{'2'},
		{'+'},
	}
	for i, tt := range tests {
		got := game_application.IsInRightSpot(tt.arg, rune(solution[0]))
		if got != expected {
			t.Errorf("test %d: expected %v but got %v", i, expected, got)
		}
	}
}

func TestIsInRightSpot_passing_arg_in_right_spot_expect_true(t *testing.T) {
	expected := true
	solution := "022+20"
	tests := []struct {
		arg rune
	}{
		{'+'},
	}
	for i, tt := range tests {
		got := game_application.IsInRightSpot(tt.arg, rune(solution[3]))
		if got != expected {
			t.Errorf("test %d: expected %v but got %v", i, expected, got)
		}
	}
}

func TestGetHints_passing_args_unavaiable_in_the_solution_expect_X_hints(t *testing.T) {
	solution := "222+20"
	expected := "XXXXXX"
	var tests = []struct {
		args string
	}{
		{"333-39"},
		{"3/3143"},
		{"8*5649"},
		{"7-35*9"},
	}
	for i, tt := range tests {
		got := game_application.GetHints(tt.args, solution)
		if got != expected {
			t.Errorf("test %d: expected %v but got %v", i, expected, got)
		}
	}
}

func TestGetHints_passing_args_avaiable_in_the_solution_but_in_the_wrong_spot_expect_T_hints(t *testing.T) {
	solution := "120/40"
	expected := []string{"XXXXXT", "XXXXTX", "XXXTXX", "XXTXXX", "XTXXXX", "TTTTTT", "TTXXXX", "TTTTXX"}
	var tests = []struct {
		args string
	}{
		{"333331"},
		{"333313"},
		{"333133"},
		{"331333"},
		{"313333"},
		{"20/104"},
		{"013*59"},
		{"4/1275"},
	}
	for i, tt := range tests {
		got := game_application.GetHints(tt.args, solution)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}

func TestInitHint_passing_args_avaiable_in_the_solution_that_is_in_the_right_spot_expect_C_hints(t *testing.T) {
	solution := "120/40"
	expected := []string{"XXXXXC", "XXXXCX", "XXXCXX", "XXCXXX", "XCXXXX", "CXXXXX"}
	var tests = []struct {
		args string
	}{
		{"333330"},
		{"333343"},
		{"333/33"},
		{"330333"},
		{"323333"},
		{"133333"},
	}
	for i, tt := range tests {
		got := game_application.GetHints(tt.args, solution)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}

func TestInitHint_simulating_a_game_expect_TCX_hints(t *testing.T) {
	solution := "120/40"
	expected := []string{"XXTXTC", "TCTXCC", "XXCXXC", "TTCCCX", "TCTTTC"}
	var tests = []struct {
		args string
	}{
		{"33/320"},
		{"/21340"},
		{"330-30"},
		{"/10/49"},
		{"42/010"},
	}
	for i, tt := range tests {
		got := game_application.GetHints(tt.args, solution)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}

func TestInitHint_passing_solution_expect_CCCCCC(t *testing.T) {
	solution := "120/40"
	expected := []string{"CCCCCC"}
	var tests = []struct {
		args string
	}{
		{"120/40"},
	}
	for i, tt := range tests {
		got := game_application.GetHints(tt.args, solution)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}
