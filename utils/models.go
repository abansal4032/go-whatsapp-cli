package utils

import (
	"flag"
	"fmt"
)

// SendArgs holds the information for sending a message.
type SendArgs struct {
	to   string
	text string
}

// ParseSendArgs parses the arguments for sending a message.
func ParseSendArgs(f *flag.FlagSet, args []string) (*SendArgs, error) {
	to := f.String("to", "", "the receipient of the message.")
	text := f.String("text", "", "the content of the message.")
	if err := f.Parse(args); err != nil {
		return nil, err
	}
	// checks for mandatory fields
	if *to == "" {
		return nil, fmt.Errorf("to field empty")
	}
	if *text == "" {
		return nil, fmt.Errorf("text field empty")
	}
	sendArgs := &SendArgs{
		to:   *to,
		text: *text,
	}
	return sendArgs, nil
}

// GetReciever gets the reciever of the message.
func (s *SendArgs) GetReciever() string {
	return s.to
}

// GetContent gets the content of the message.
func (s *SendArgs) GetContent() string {
	return s.text
}
