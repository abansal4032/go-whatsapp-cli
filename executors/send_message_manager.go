package executors

import (
	"fmt"
	"github.com/abansal4032/go-whatsapp-cli/utils"
	"os"
	"strings"
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
		return textMessage(msgMetadata)
	case utils.IMAGEMESSAGEKEY:
		return imageMessage(msgMetadata)
	case utils.VIDEOMESSAGEKEY:
		return videoMessage(msgMetadata)
	case utils.AUDIOMESSAGEKEY:
		return audioMessage(msgMetadata)
	case utils.MEDIAMESSAGEKEY:
		return mediaMessage(msgMetadata)
	default:
		return nil, fmt.Errorf("unknown message type : %v\n feel free to raise an issue on the github repo with the desired message type to send", msgType)
	}
}

func textMessage(msgMetadata interface{}) (interface{}, error) {
	sendTextArgs, ok := msgMetadata.(*utils.SendTextArgs)
	if !ok {
		return nil, fmt.Errorf("cannot read args for text message")
	}
	text := sendTextArgs.Content() + utils.MESSAGEFOOTER
	receiver := sendTextArgs.Receiver()
	msg := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: receiver + utils.CONTACTSUFFIX,
		},
		Text: text,
	}
	return msg, nil
}

func videoMessage(msgMetadata interface{}) (interface{}, error) {
	sendMediaArgs, ok := msgMetadata.(*utils.SendMediaArgs)
	if !ok {
		return nil, fmt.Errorf("cannot read args for media message")
	}
	text := sendMediaArgs.Content() + utils.MESSAGEFOOTER
	receiver := sendMediaArgs.Receiver()
	contentType := sendMediaArgs.ContentType()
	media, _ := os.Open(sendMediaArgs.Filepath())
	msg := whatsapp.VideoMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: receiver + utils.CONTACTSUFFIX,
		},
		Type:    contentType,
		Caption: text,
		Content: media,
	}
	return msg, nil
}

func imageMessage(msgMetadata interface{}) (interface{}, error) {
	sendMediaArgs, ok := msgMetadata.(*utils.SendMediaArgs)
	if !ok {
		return nil, fmt.Errorf("cannot read args for media message")
	}
	text := sendMediaArgs.Content() + utils.MESSAGEFOOTER
	receiver := sendMediaArgs.Receiver()
	contentType := sendMediaArgs.ContentType()
	media, _ := os.Open(sendMediaArgs.Filepath())
	msg := whatsapp.ImageMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: receiver + utils.CONTACTSUFFIX,
		},
		Type:    contentType,
		Caption: text,
		Content: media,
	}
	return msg, nil
}

func audioMessage(msgMetadata interface{}) (interface{}, error) {
	sendMediaArgs, ok := msgMetadata.(*utils.SendMediaArgs)
	if !ok {
		return nil, fmt.Errorf("cannot read args for media message")
	}
	receiver := sendMediaArgs.Receiver()
	contentType := sendMediaArgs.ContentType()
	media, _ := os.Open(sendMediaArgs.Filepath())
	msg := whatsapp.AudioMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: receiver + utils.CONTACTSUFFIX,
		},
		Type:    contentType,
		Content: media,
	}
	return msg, nil
}

func mediaMessage(msgMetadata interface{}) (interface{}, error) {
	sendMediaArgs, ok := msgMetadata.(*utils.SendMediaArgs)
	if !ok {
		return nil, fmt.Errorf("cannot read args for media message")
	}
	contentType := sendMediaArgs.ContentType()
	t := strings.Split(contentType, "/")[0]
	switch t {
	case "image":
		return imageMessage(msgMetadata)
	case "video":
		return videoMessage(msgMetadata)
	case "audio":
		return audioMessage(msgMetadata)
	default:
		return nil, fmt.Errorf("unknown content type : %v\n feel free to raise an issue with the content type on the github repo", contentType)
	}
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
