#!/bin/bash -e

go build -ldflags="-w -s" -o deploy cmd/watermark/watermark.go
