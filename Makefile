# 检测操作系统并设置变量
IS_WINDOWS := $(if $(findstring Windows_NT,$(OS)),yes,no)

ifeq ($(IS_WINDOWS),yes)
BUILD_ADMIN = powershell -ExecutionPolicy Bypass -File scripts/admin/build.ps1
BUILD_ADMIN_RELEASE = powershell -ExecutionPolicy Bypass -File scripts/admin/build.ps1 release

else
BUILD_ADMIN = sh scripts/admin/build.sh
BUILD_ADMIN_RELEASE = sh scripts/admin/build.sh release

endif

build_admin:
	@$(BUILD_ADMIN)

build_admin_release:
	@$(BUILD_ADMIN_RELEASE)