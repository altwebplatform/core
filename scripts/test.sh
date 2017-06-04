#!/usr/bin/env bash

go test -race -timeout 45s $(glide novendor)
