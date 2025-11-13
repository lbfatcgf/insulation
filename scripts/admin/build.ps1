# PowerShell build script for admin service

# Exit on error
$ErrorActionPreference = "Stop"
# 检查并安装 swag
if (-not (Get-Command `swag` -ErrorAction SilentlyContinue)) {
   
    go install github.com/swaggo/swag/cmd/swag@latest
    # 刷新环境变量确保新安装的命令可用
}
swag fmt
# Generate Swagger documentation
swag init -g cmd\admin\main.go -o apis\admin

# Get current directory
$dir = Get-Location

# Build Go application
go build -o "$dir\build\output\admin.run" "$dir\cmd\admin"