@echo off

if exist devtool.exe (
	devtool.exe %*
) else (
	go run devtool/main.go %*
)