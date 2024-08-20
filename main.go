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
type Handler interface {
	Handle(inputs []string, output io.Writer, calculator Calculator) error
}

type calcHandler struct{}

func (calc calcHandler) Handle(inputs []string, output io.Writer, calculator Calculator) error {
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
	result := calculator.Calculate(addend1, addend2)

	if _, err := fmt.Fprintln(output, result); err != nil {
		return err
	}
	return nil
}

func main() {
	var (
		inputs     []string   = os.Args[1:]
		calculator Calculator = calc.Addition{}
		output     io.Writer  = os.Stdout
		handler    Handler    = &calcHandler{}
	)
	err := handler.Handle(inputs, output, calculator)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
