package executors

import (
	"fmt"
	"github.com/Rhymen/go-whatsapp"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type textHandler struct{}

// HandleError implements the handler interface for go-whatsapp
func (t textHandler) HandleError(err error) {
	// TODO : handle go routine here
	fmt.Printf("error in textHandler : %v", err)
	return
}

// checks if the string is of interest to the handler.
// checks for "youtube:" tag in this case.
// returns the "song:" tag if successful.
func interesting(str string) (string, bool) {
	parts := strings.Split(str, ":")
	if len(parts) < 2 {
		return "", false
	}
	fmt.Println(parts, strings.EqualFold(strings.TrimSpace(parts[0]), "youtube"), strings.TrimSpace(parts[0]))
	if strings.EqualFold(strings.TrimSpace(parts[0]), "youtube") {
		return strings.TrimSpace(parts[1]), true
	}
	return "", false
}

// HandleTextMessage implements the text message handler interface for go-whatsapp
func (t textHandler) HandleTextMessage(msg whatsapp.TextMessage) {
	fmt.Println("received message")
	text := msg.Text
	if song, ok := interesting(text); ok && song != "" {
		fmt.Printf("going to play %v on youtube\n", song)
	}
}

// StartTextReceiver starts the handler for the text messages received
func StartTextReceiver() error {
	wac, err := whatsapp.NewConn(5 * time.Second)
	if err != nil {
		return fmt.Errorf("error creating connection: %v\n", err)
	}

	err = LoginWithConnection(wac)
	if err != nil {
		return fmt.Errorf("error logging in: %v\n", err)
	}

	<-time.After(3 * time.Second)

	handler := textHandler{}
	wac.AddHandler(handler)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	fmt.Println("closing the receiver")
	wac.Disconnect()
	return nil
}
