#!/usr/bin/env bash
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
go build -ldflags "-s -w" -o "$SCRIPT_DIR/bin/safira" "$SCRIPT_DIR/src/safira.go" 
