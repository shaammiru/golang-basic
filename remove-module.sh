#!/bin/bash

if grep -q "$1" "go.work"; then
  sed -i "\@$1@d" go.work
  rm -rf $1
  echo "Module $1 removed"
else
  echo "Module $1 not found"
fi