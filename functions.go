package genqlfinance

import (
	"fmt"
	"math"
)

//	Calculates extact amount to be collected based on a transaction fee
//
// --------------------------------------------------
// | index |    type    |       description         |
// |-------|------------|---------------------------|
// |   0   |   float64  |     original amount       |
// |   1   |   float64  |     transaction fee       |
// --------------------------------------------------
func AdjustFunc(args []any) (any, error) {
	err := Guard(2, args)
	if err != nil {
		return nil, err
	}

	originalAmount, ok := args[0].(float64)
	if !ok {
		return nil, fmt.Errorf("expected float64 but found %T", args[0])
	}

	transactionFee, ok := args[1].(float64)
	if !ok {
		return nil, fmt.Errorf("expected float64 but found %T", args[0])
	}
	if transactionFee <= 0 {
		return nil, fmt.Errorf("invalid transaction fee")
	}

	return math.Ceil(originalAmount + ((transactionFee * originalAmount) / (1 - transactionFee))), nil
}

//	Adds commission to an amount
//
// --------------------------------------------------
// | index |    type    |       description         |
// |-------|------------|---------------------------|
// |   0   |   float64  |     original amount       |
// |   1   |   float64  |        commission         |
// --------------------------------------------------
func ApplyCommisionFunc(args []any) (any, error) {
	err := Guard(2, args)
	if err != nil {
		return nil, err
	}

	originalAmount, ok := args[0].(float64)
	if !ok {
		return nil, fmt.Errorf("expected float64 but found %T", args[0])
	}

	commission, ok := args[1].(float64)
	if !ok {
		return nil, fmt.Errorf("expected float64 but found %T", args[0])
	}
	if commission <= 0 {
		return nil, fmt.Errorf("invalid comission fee")
	}

	return math.Ceil(originalAmount + (originalAmount * commission / 100)), nil
}

func Guard(n int, args []any) error {
	if len(args) < n {
		return fmt.Errorf("too few arguments")
	}
	if len(args) > n {
		return fmt.Errorf("too many arguments")
	}
	return nil
}

func Export() map[string]func([]any) (any, error) {
	return map[string]func([]any) (any, error){
		"adjust":          AdjustFunc,
		"applycommission": ApplyCommisionFunc,
	}
}
