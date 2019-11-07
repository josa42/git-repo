#!/bin/bash

if [[ "$1" == "" ]]; then
  exit 1
fi

sed -E -e "s/\"Git Repo [^\"]+\"/\"Git Repo $1\"/" -i.bak main.go || exit 1
rm main.go.bak || exit 1
