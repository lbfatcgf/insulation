#!/bin/bash
set -e -u
swag fmt
swag init -g cmd/admin/main.go -o apis/admin

dir=$(pwd)

go build -o "$dir/build/output/admin.run" "$dir/cmd/admin"