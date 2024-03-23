package wordy

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Operation string

const (
	Addition       Operation = "+"
	Division       Operation = "/"
	Multiplication Operation = "*"
	Subtraction    Operation = "-"
)

func lex(question string) ([]interface{}, error) {
	words := strings.Fields(strings.TrimSuffix(strings.TrimPrefix(question, "What is"), "?"))

	var tokens []interface{}
	for i := range words {
		number, err := strconv.Atoi(words[i])

		switch {
		case err == nil:
			tokens = append(tokens, number)
		case words[i] == "by":
			// SKIP
		case words[i] == "divided" && words[i+1] == "by":
			tokens = append(tokens, Division)
		case words[i] == "minus":
			tokens = append(tokens, Subtraction)
		case words[i] == "multiplied" && words[i+1] == "by":
			tokens = append(tokens, Multiplication)
		case words[i] == "plus":
			tokens = append(tokens, Addition)
		default:
			return nil, fmt.Errorf("invalid operation: %s", words[i])
		}
	}

	if len(tokens)%2 == 0 {
		return nil, errors.New("operators missing operands")
	}

	return tokens, nil
}

func interpret(tokens []interface{}) (int, bool) {
	first, remaining := tokens[0], tokens[1:]

	accumulator, ok := first.(int)
	if !ok {
		return 0, false
	}

	for cursor := 0; cursor < len(remaining)-1; cursor += 2 {
		operator, ok := remaining[cursor].(Operation)
		if !ok {
			return 0, false
		}

		operand, ok := remaining[cursor+1].(int)
		if !ok {
			return 0, false
		}

		switch operator {
		case Addition:
			accumulator += operand
		case Division:
			accumulator /= operand
		case Multiplication:
			accumulator *= operand
		case Subtraction:
			accumulator -= operand
		}
	}

	return accumulator, true
}

func Answer(question string) (int, bool) {
	tokens, err := lex(question)
	if err != nil {
		return 0, false
	}

	return interpret(tokens)
}
