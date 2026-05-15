#!/bin/bash

echo "BUF DEP UPDATE"
buf dep update

echo "BUF LINT"
buf lint

echo "BUF GEN"
buf generate
