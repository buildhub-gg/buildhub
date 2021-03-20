#!/bin/bash

go get -v -t -d ./...
go get github.com/bazelbuild/bazelisk
go get github.com/google/wire/cmd/wire
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega/...