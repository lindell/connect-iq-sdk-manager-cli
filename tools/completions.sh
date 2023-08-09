#!/bin/sh
set -e
rm -rf completions
mkdir completions
for sh in bash zsh fish; do
	go run main.go completion "$sh" >"completions/connect-iq-sdk-manager.$sh"
done
