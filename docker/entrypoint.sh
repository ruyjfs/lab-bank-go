#!/bin/sh

echo "Clean cache..."
go clean -modcache

echo "Install..."
go mod init github.com/ruyjfs/lab-bank-go

# echo "Generate ..."
# # go run github.com/99designs/gqlgen init
# gqlgen generate

exec "$@"
