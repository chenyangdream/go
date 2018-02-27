package common

import (
	"errors"
	"time"
)

type Args struct {
	A, B float32
}

type Result struct {
	Value float32
}

type MathService struct {
}

func (s *MathService) Add(args *Args, result *Result) error {
	result.Value = args.A + args.B
	time.Sleep(10 * time.Second)
	return nil
}

func (s *MathService) Divide(args *Args, result *Result) error {
	if args.B == 0 {
		return errors.New("除数不能为零!")
	}

	result.Value = args.A / args.B
	time.Sleep(10 * time.Second)
	return nil
}
