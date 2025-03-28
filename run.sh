#!/bin/bash

if ! [ -f private.pem ]; then
    echo "Generate private key"
    openssl genrsa -out private.pem 2048
fi

go run ./cmd/app
