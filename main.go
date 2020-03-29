package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/abansal4032/go-whatsapp-cli/executors"
	"github.com/abansal4032/go-whatsapp-cli/utils"
)

func commandUsage(err error, f *flag.FlagSet) {
	fmt.Println("correct command usage is described below:")
	f.PrintDefaults()
}

func errorHandler(err error) {
	// TODO : graceful error handling here
	if err != nil {
		fmt.Printf("flow errored out : %v\n", err.Error())
	}
}

func main() {
	sendTextCmd := flag.NewFlagSet("sendText", flag.ExitOnError)
	var err error
	defer func() {
		errorHandler(err)
	}()

	switch os.Args[1] {

	case "login":
		if err = executors.Login(); err != nil {
			return
		}

	case "logout":
		if err = executors.Logout(); err != nil {
			return
		}

	case "sendText":
		var args *utils.SendTextArgs
		args, err = utils.ParseSendArgs(sendTextCmd, os.Args[2:])
		if err != nil {
			commandUsage(err, sendTextCmd)
			return
		}
		if err = executors.SendMessage(args, utils.TEXTMESSAGEKEY); err != nil {
			return
		}

	default:
		// TODO : add a list of all permitted actions
		err = fmt.Errorf("wrong command provided. please see below for the list of permitted actions")
		return
	}
}
