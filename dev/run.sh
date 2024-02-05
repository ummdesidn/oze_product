#!/usr/bin/env bash

docker build . -t go-001
docker run -i -t -p 9000:8080 go-001