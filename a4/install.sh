#!/bin/bash

## =================================================================
## A4MD INSTALLATION
##
## This script will install all the packages that are needed to
## build and run the a4md Tool.
##
## Supported environments:
##  * Ubuntu 18.04
##  * macOS
## =================================================================

dir=~/Library/a4md

if [[ ! -e ~/Library ]]; then
    mkdir ~/Library
fi

## build a4md
go mod tidy
go build a4md.go

if [[ -e /usr/local/bin/a4md ]]; then 
    sudo rm -f /usr/local/bin/a4md
fi
sudo cp a4md /usr/local/bin/a4md

g++ -std=c++17 main.cpp -o ~/Library/a4md/a4filter

if [[ ! -e $dir ]]; then
    mkdir $dir
elif [[ ! -d $dir ]]; then
    echo "$dir already exists but is not a directory" 1>&2
fi

sudo cp article1.html ~/Library/a4md/article1.html
sudo cp article2.html ~/Library/a4md/article2.html
sudo cp article3.html ~/Library/a4md/article3.html
sudo cp article4.html ~/Library/a4md/article4.html

