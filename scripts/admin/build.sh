#!/bin/bash
set -e -u

if ! command -v `swag` &> /dev/null; then
    go install github.com/swaggo/swag/cmd/swag@latest
fi
swag fmt
swag init -g cmd/admin/main.go -o apis/admin

dir=$(pwd)

go build -o "$dir/build/output/admin.run" "$dir/cmd/admin"