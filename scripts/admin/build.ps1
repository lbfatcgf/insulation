# PowerShell build script for admin service

# Exit on error
$ErrorActionPreference = "Stop"

swag fmt
# Generate Swagger documentation
swag init -g cmd\admin\main.go -o apis\admin

# Get current directory
$dir = Get-Location

# Build Go application
go build -o "$dir\build\output\admin.run" "$dir\cmd\admin"