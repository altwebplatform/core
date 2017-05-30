#!/usr/bin/env bash

cd reloader
env GOOS=linux GOARCH=386 go build -v -o ../bin/reloader
cd ../

env GOOS=linux GOARCH=386 go build -v -o app
