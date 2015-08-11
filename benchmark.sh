#!/bin/sh

time go test -v -benchmem -bench . | tee benchmarks.txt
