#!/bin/bash

# Create directory
mkdir $1

# Create module
cd $1
go mod init $1

# Create main.go
touch main.go

# Add package main
echo "package main" >> main.go

# Add to workspace
cd ..
go work use $1

echo "Module $1 created"