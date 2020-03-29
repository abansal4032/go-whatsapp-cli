package main

import (
	"fmt"
	"flag"
	"os"

	"github.com/abansal4032/go-whatsapp-cli/executors"
	"github.com/abansal4032/go-whatsapp-cli/utils"
)

func commandUsage(err error) {
	// TODO : graceful error handling here
	if err != nil {
		fmt.Println("correct command usage is described below:")
		flag.PrintDefaults()
	}
}

func main() {
	sendTextCmd := flag.NewFlagSet("sendText", flag.ExitOnError)
	var err error
	defer func() {
		commandUsage(err)
	}()

	switch os.Args[1] {
	
	case "login":
		if err = executors.Login(); err != nil {
			fmt.Println(err)
			return
		}

	case "logout":
		if err = executors.Logout(); err != nil {
			fmt.Println(err)
			return
		}

	case "sendText":
		//var args *utils.SendArgs
		_, err = utils.ParseSendArgs(sendTextCmd, os.Args[2:])
		if err != nil {
			fmt.Println(err)
			return
		}
		/*if err = executors.SendText(args); err != nil {
			fmt.Println(err)
			return
		}*/
	
	default:
		fmt.Println("wrong command provided. please see below for the list of permitted actions")
		flag.PrintDefaults()
		return
	
	}
}