#!/bin/bash

rm -rf dolt_*.auto.go dolt_*auto_test.go dolt_test_util.go

go build

./dolt -generate

go build
