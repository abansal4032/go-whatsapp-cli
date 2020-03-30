package utils

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

type commonArgs struct {
	to   string
	text string
}

// SendTextArgs holds the information for sending a text message.
type SendTextArgs struct {
	commonArgs
}

// SendMediaArgs holds the information for sending a media message(file).
type SendMediaArgs struct {
	commonArgs
	filepath    string
	contentType string
}

// ParseSendTextArgs parses the arguments for sending a text message.
func ParseSendTextArgs(f *flag.FlagSet, args []string) (*SendTextArgs, error) {
	to := f.String("to", "", "the recipient of the message.")
	text := f.String("text", "", "the content of the message.")
	if err := f.Parse(args); err != nil {
		return nil, err
	}
	sendArgs := &SendTextArgs{
		commonArgs{
			to:   *to,
			text: *text,
		},
	}
	if err := sendArgs.Validate(); err != nil {
		return nil, err
	}
	return sendArgs, nil
}

// Validate validates the body.
func (s *SendTextArgs) Validate() error {
	return s.validate()
}

// ParseSendMediaArgs parses the arguments for sending a media message(file).
func ParseSendMediaArgs(f *flag.FlagSet, args []string) (*SendMediaArgs, error) {
	to := f.String("to", "", "the recipient of the message.")
	text := f.String("text", "", "the content of the message.")
	filepath := f.String("filepath", "", "the absolute of the file to be sent.")
	if err := f.Parse(args); err != nil {
		return nil, err
	}
	sendArgs := &SendMediaArgs{
		commonArgs: commonArgs{
			to:   *to,
			text: *text,
		},
		filepath: *filepath,
	}
	sendArgs.contentType = getFileContentType(*filepath)
	if err := sendArgs.Validate(); err != nil {
		return nil, err
	}
	return sendArgs, nil
}

// Validate validates the body.
func (s *SendMediaArgs) Validate() error {
	// TODO : Do not need text for audioMessage
	if err := s.validate(); err != nil {
		return err
	}
	if s.filepath == "" || !exists(s.filepath) {
		return fmt.Errorf("error validating args : the filepath %v does not exist", s.filepath)
	}
	if s.contentType == "" {
		return fmt.Errorf("error validating args : could not get the content type of file")
	}
	return nil
}

// Filepath returns the filepath specified by the user.
func (s *SendMediaArgs) Filepath() string {
	return s.filepath
}

// ContentType returns the mime-type of the file content.
func (s *SendMediaArgs) ContentType() string {
	return s.contentType
}

func exists(filepath string) bool {
	if _, err := os.Stat(filepath); err != nil {
		// TODO : should check for IsNotExist for error here
		return false
	}
	return true
}

// Receiver gets the receiver of the message.
func (c *commonArgs) Receiver() string {
	return c.to
}

// Content gets the content of the message.
func (c *commonArgs) Content() string {
	return c.text
}

func (c *commonArgs) validate() error {
	if c.to == "" {
		return fmt.Errorf("error validating args : empty receiver")
	}
	return nil
}

func getFileContentType(filepath string) string {
	f, _ := os.Open(filepath)
	buffer := make([]byte, 512)
	_, err := f.Read(buffer)
	if err != nil {
		return ""
	}
	contentType := http.DetectContentType(buffer)
	return contentType
}
