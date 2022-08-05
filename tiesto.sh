#!/bin/bash


program=$(echo '.[] |select(.hometown|contains(@@)) // select(.residence|contains(@@))|.name' | sed -e s/@@/\"$1\"/g)
go run cmd/tiesto/main.go | jq "$program"
