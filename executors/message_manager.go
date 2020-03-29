package executors

import (
	"fmt"
	"time"

	"github.com/Rhymen/go-whatsapp"
)

// SendText sends the provided text message to the receipient.
func SendText(text, reciever string) error {
	wac, err := whatsapp.NewConn(5 * time.Second)
	if err != nil {
		return fmt.Errorf("error creating connection: %v\n", err)
	}

	err = LoginWithConnection(wac)
	if err != nil {
		return fmt.Errorf("error logging in: %v\n", err)
	}

	<-time.After(3 * time.Second)

	text += "\n sent using github.com/abansal4032/go-whatsapp-cli"
	msg := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: reciever + "@s.whatsapp.net",
		},
		Text: text,
	}

	msgID, err := wac.Send(msg)
	if err != nil {
		return fmt.Errorf("error sending message: %v", err)
	}

	fmt.Printf("successfully sent, messaageID : %v\n", msgID)
	return nil
}
