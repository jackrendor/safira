SET SCRIPT_DIR=%~dp0
go build -ldflags "-s -w" -o "%SCRIPT_DIR%/bin/safira" "%SCRIPT_DIR%/src/safira.go"
