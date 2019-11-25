#!/bin/bash

if [[ "$1" == "" ]]; then
  exit 1
fi

sed -E -e "s/(Version: *)\"[^\"]+\"/\1\"$1\"/" -i.bak cmd/root.go || exit 1
rm -f cmd/root.go.bak || exit 1

