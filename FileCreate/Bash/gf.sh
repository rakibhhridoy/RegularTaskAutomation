#!/bin/bash

touch $1.go
echo "package main" > $1.go
echo " " >> $1.go 
echo "import 'fmt'" >> $1.go
echo " " >> $1.go
echo "func main() {}" >> $1.go
vim $1.go
