#!/bin/bash
set -e -u

tag=$0
if [ -z "${tag+x}" ]; then
    tag=dev
elif [ "$tag" != "release" ]; then
    tag=dev
fi
if ! command -v `swag` &> /dev/null; then
    go install github.com/swaggo/swag/cmd/swag@latest
fi
swag fmt
swag init -g cmd/admin/main.go -o apis

dir=$(pwd)


if [ "$tag" = "release" ]; then
    echo "编译环境：release"
    go build -ldflags="-s -w" -o "$dir/build/output/admin.run" "$dir/cmd/admin" 
    exit 0
fi

if [ "$tag" = "dev" ]; then
    echo "编译环境：dev"
    go build -o "$dir/build/output/admin.run" "$dir/cmd/admin"
    exit 0
fi


