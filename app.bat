@echo off

if exist bin\app.exe (
	bin\app.exe %*
) else (
	go run main.go %*
)