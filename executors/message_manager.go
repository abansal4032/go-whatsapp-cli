package executors

import (
	"fmt"
	"github.com/abansal4032/go-whatsapp-cli/utils"
	"time"

	"github.com/Rhymen/go-whatsapp"
)

// SendMessage sends the provided message to the recipient.
func SendMessage(args interface{}, msgType string) error {
	msg, err := composeMessage(args, msgType)
	if err != nil {
		return fmt.Errorf("error while composing message: %v\n", err)
	}
	return sendMessage(msg)
}

func composeMessage(msgMetadata interface{}, msgType string) (interface{}, error) {
	switch msgType {
	case utils.TEXTMESSAGEKEY:
		sendTextArgs, ok := msgMetadata.(*utils.SendTextArgs)
		if !ok {
			return nil, fmt.Errorf("cannot read args for text message")
		}
		text := sendTextArgs.GetContent() + utils.MESSAGEFOOTER
		receiver := sendTextArgs.GetReceiver()
		msg := whatsapp.TextMessage{
			Info: whatsapp.MessageInfo{
				RemoteJid: receiver + utils.CONTACTSUFFIX,
			},
			Text: text,
		}
		return msg, nil
	case utils.IMAGEMESSAGEKEY:
	case utils.VIDEOMESSAGEKEY:
	default:
		return nil, fmt.Errorf("unknown message type : %v\n", msgType)
	}
	// Gotta keep the compiler happy
	return nil, fmt.Errorf("unknown message type : %v\n", msgType)
}

func sendMessage(msg interface{}) error {
	wac, err := whatsapp.NewConn(5 * time.Second)
	if err != nil {
		return fmt.Errorf("error creating connection: %v\n", err)
	}

	err = LoginWithConnection(wac)
	if err != nil {
		return fmt.Errorf("error logging in: %v\n", err)
	}

	<-time.After(3 * time.Second)

	msgID, err := wac.Send(msg)
	if err != nil {
		return fmt.Errorf("error sending message: %v", err)
	}

	fmt.Printf("successfully sent, messaageID : %v\n", msgID)
	return nil
}
