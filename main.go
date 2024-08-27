package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/mdwhatcott/calcy-lib/calc"
)

type Calculator interface {
	Calculate(a, b int) int
}

type Handler struct {
	output     io.Writer
	calculator Calculator
}

func NewHandler(calc Calculator, output io.Writer) *Handler {
	return &Handler{output, calc}
}

func (this *Handler) Handle(inputs []string) error {
	if len(inputs) != 2 {
		return errors.New("usage: go run main.go <a> <b>")
	}
	addend1, err := strconv.Atoi(inputs[0])
	if err != nil {
		return errors.New("failed to parse parameter 1")
	}
	addend2, err := strconv.Atoi(inputs[1])
	if err != nil {
		return errors.New("failed to parse parameter 2")
	}
	result := this.calculator.Calculate(addend1, addend2)

	if _, err := fmt.Fprintln(this.output, result); err != nil {
		return err
	}
	return nil
}

func main() {
	var (
		inputs     []string   = os.Args[1:]
		calculator Calculator = calc.Addition{}
		output     io.Writer  = os.Stdout
	)
	handler := NewHandler(calculator, output)

	err := handler.Handle(inputs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
