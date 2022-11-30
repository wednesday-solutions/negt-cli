#!/usr/bin/env bash

go test ./... -gcflags=all=-l -coverprofile=coverage.out 