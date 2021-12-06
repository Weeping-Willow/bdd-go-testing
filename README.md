## Description

Small exercise on writing bdd and acceptance tests

This was made so i can try out something new, so the structure is kinda wonky

## Requirements

- golang 
- maybe ginkgo, not sure if tests run locally without it(https://github.com/onsi/ginkgo)
- github.com/vakenbolt/go-test-report/ this if you want fancy html output

if you have golang and make you can run `make dep` to install the bottom two requirements

## Commands

- `make test` just runs tests
- `make test-fancy` run ginkgo tests with their fancy output
- `make test-html` outputs tests pass and fails in a html file