# PowerShell build script for admin service

# Exit on error
$ErrorActionPreference = "Stop"

# 初始化变量 $tag
$tag = $args[0]
if (-not $tag) {
    $tag = "dev"
} elseif ($tag -ne "release") {
    $tag = "dev"
}
# 检查并安装 swag
if (-not (Get-Command swag  -ErrorAction SilentlyContinue)) {
    go install "github.com/swaggo/swag/cmd/swag@latest"
}
swag fmt
# Generate Swagger documentation
swag init -g cmd\admin\main.go -o apis

# Get current directory
$dir = Get-Location

# Build Go application
# 根据 $tag 的值执行不同的编译逻辑
if ($tag -eq "release") {
    Write-Output "编译环境：release"
    go build -ldflags="-s -w" -o "$dir/build/output/admin.run" "$dir/cmd/admin"
    exit 0
}

if ($tag -eq "dev") {
    Write-Output "编译环境：dev"
    go build -o "$dir/build/output/admin.run" "$dir/cmd/admin"
    exit 0
}