#!/bin/sh

# if app exists, run it
if [ -f "bin/app" ]; then
	./bin/app $*
else
	go run main.go $*
fi
