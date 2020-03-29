# go-whatsapp-client
This repo builds over Rhymen/go-whatsapp to provide a cli for whatsapp.


# Running

The dependencies will be vendored in the repository. Due to difference in how go 1.13 and go 1.14 treats vendor directory, there are two ways to run depending on the version:
**go 1.13** : GO111MODULE=off go run main.go
**go 1.14** : go run main.go
