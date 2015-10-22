#!/bin/bash

platforms=( darwin linux windows )

archs=( 386 amd64 )

for i in "${!platforms[@]}"
do
  for j in "${!archs[@]}"
  do
    echo ${platforms[$i]}/${archs[$j]}
    GOOS=${platforms[$i]} GOARCH=${archs[$j]} go build -o rethinkdb-backerupperer-${platforms[$i]}-${archs[$j]} rethinkdb-backerupperer.go
  done
done
