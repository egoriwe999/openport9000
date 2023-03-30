# Multi Stage Build
Build docker container on Golang project (Multi stage build)

Images : golang:alpine and scratch


How do up container?

docker build -t openport:latest -f /openport9000/DockerFile .

We have image , after we can run container

docker run --name golangport openport:latest

Completed!
