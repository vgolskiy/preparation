# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

# Working directory for application
WORKDIR /mini_webserv

# Getting nessesary golang modules
COPY go.mod ./
COPY go.sum ./

# Installing modules
RUN go mod download

# Copying application files to docker
COPY *.go ./

# Application compilation
RUN go build -o /mini_webserv

# Port to call application inside of docker
EXPOSE 8888

# Starting application
CMD [ "/mini_webserv" ]
