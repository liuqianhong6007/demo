#!/bin/bash

rm -rf dolt_*.auto.go dolt_*auto_test.go dolt_test_util.go dolt_api_doc.md

go build

./dolt -generate

go fmt
go build
