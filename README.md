# go-whatsapp-client
This repo builds over Rhymen/go-whatsapp to provide a cli for whatsapp.

# Building
The Makefile consists of rules to help you build the project. 

To build the project including the linters and fmt:\
**make all**

To only build the binary:\
**make build**

To clean:\
**make clean**

The dependencies are vendored in the repository. Due to difference in how go 1.13 and go 1.14 treats vendor directory, there are two ways to run depending on the version:\
**go 1.13** : GO111MODULE=off go build\
**go 1.14** : go build\
Make sure you make the changes in the Makefile if you are using go 1.13. The current file is amied for go 1.14.

# Running
Current commands installed:\
**1. whatsapp-cli login**: Logs you in\
**2. whatsapp-cli logout**: Logs you out\
**3. whatsapp-cli sendText --to <phone_number> --text <text_message>**: Sends the text_message to the phone_number

**NOTE** : More coming soon...