package utils

import (
	"flag"
	"fmt"
)

// SendTextArgs holds the information for sending a message.
type SendTextArgs struct {
	to   string
	text string
}

// ParseSendArgs parses the arguments for sending a message.
func ParseSendArgs(f *flag.FlagSet, args []string) (*SendTextArgs, error) {
	to := f.String("to", "", "the recipient of the message.")
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
	sendArgs := &SendTextArgs{
		to:   *to,
		text: *text,
	}
	return sendArgs, nil
}

// GetReceiver gets the receiver of the message.
func (s *SendTextArgs) GetReceiver() string {
	return s.to
}

// GetContent gets the content of the message.
func (s *SendTextArgs) GetContent() string {
	return s.text
}
