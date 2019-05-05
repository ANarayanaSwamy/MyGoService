@echo off
echo "Started to compile"
set GOPATH=%CD%
go install service
echo "Executing...."
bin\service