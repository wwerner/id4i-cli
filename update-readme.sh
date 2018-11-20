#!/usr/bin/env bash

set -o errexit  # exit on error
set -o nounset  # don't allow unset variables
# set -o xtrace # enable for debugging

echo "----" > help.txt
go run main.go help >> help.txt
echo "----" >> help.txt

echo "----" >> help.txt
go run main.go help guids >> help.txt
echo "----" >> help.txt

echo "----" >> help.txt
go run main.go help storage >> help.txt
echo "----" >> help.txt

echo "----" >> help.txt
go run main.go help history >> help.txt
echo "----" >> help.txt

echo ":WARNING - GENERATED FILE - DO NOT EDIT THIS FILE DIRECTLY, USE THE TEMPLATE AND update-readme.sh:" > README.adoc
sed '/=== Help/r help.txt' README.tmpl.adoc >> README.adoc