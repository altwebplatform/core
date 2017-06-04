#!/usr/bin/env bash

go fmt $(glide novendor)
go vet $(glide novendor)
