## Description

Small exercise on writing bdd and acceptance tests, i wrote 4 different scenarios that test https://gorest.co.in/
- Create a user. Positive scenario
- Update user name. Positive scenario
- Create a user with email already in use. Negative scenario.
- Remove the already removed user. Negative scenario.

This was made so i can try out something new, so the structure is kinda wonky, but it can still work, if you were to expand the project

## Requirements

- golang 
- maybe ginkgo, not sure if tests run locally without it(https://github.com/onsi/ginkgo)
- github.com/vakenbolt/go-test-report/ this if you want fancy html output

if you have golang and make, you can run `make dep` to install the bottom two requirements

## Commands

- `make test` just runs tests
- `make test-fancy` run ginkgo tests with their fancy output
- `make test-html` outputs tests pass and fails in a html file

## Notes

- As a test suite is running from a single Test, all of the scenarios are counted as one test in the html output tool, you could probably mitigate by using some weird struture, or implementing your own logger, but it's a bit too much work for this
- Honestly i still prefer just writing regular unit tests or integration tests, but i could see places where this is usefull
