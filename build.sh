#!/bin/bash
go build -o out/ ./cmd/reader
go build -o out/ ./cmd/store
chmod +x out/*