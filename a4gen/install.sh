#!/bin/bash

## =================================================================
## A4GEN INSTALLATION
##
## This script will install all the packages that are needed to
## build and run the a4gen Tool.
##
## Supported environments:
##  * Ubuntu 18.04
##  * macOS
## =================================================================


if [[ ! -e ~/Library ]]; then
    mkdir ~/Library
fi

## build a4md
go mod tidy
go build a4gen.go

if [[ -e /usr/local/bin/a4gen ]]; then 
    sudo rm -f /usr/local/bin/a4gen
fi
sudo cp a4gen /usr/local/bin/a4gen
