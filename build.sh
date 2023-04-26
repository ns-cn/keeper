#!/bin/zsh
rm -f releases/*
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o ./releases/keeper_linux .
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" -o ./releases/keeper.exe .
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-w -s" -o ./releases/keeper_darwin .