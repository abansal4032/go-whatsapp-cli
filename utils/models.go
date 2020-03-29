package utils

import (
	"flag"
	"fmt"
)

type SendArgs struct {
	to	string
}

func ParseSendArgs(f *flag.FlagSet, args []string) (*SendArgs, error) {
	to := f.String("to", "", "the receipient of the message.")
	if err := f.Parse(args); err != nil {
		return nil, err
	}
	// checks for mandatory fields
	if *to == "" {
		return nil, fmt.Errorf("wrong arguments provided")
	}
	sendArgs := &SendArgs{
		to: *to,
	}
	return sendArgs, nil
}
