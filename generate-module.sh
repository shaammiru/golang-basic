#!/bin/bash

mkdir $1
cd $1
go mod init $1
touch main.go

echo "package main" >> main.go
echo "" >> main.go
echo "func main() {" >> main.go
echo "" >> main.go
echo "}" >> main.go

cd ..
go work use $1

echo "Module $1 created"