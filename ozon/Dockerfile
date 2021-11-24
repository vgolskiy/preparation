# Container for application compilation
FROM golang:1.12.0-alpine3.9 as builder

RUN apk add git

LABEL maintainer="<>"

RUN mkdir /ozon
ADD . /ozon
WORKDIR /ozon

COPY go.mod go.sum ./

# Download all the dependencies
RUN go mod download

# Copying application files to docker
COPY ./src ./

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Build the Go app
RUN go build -a -installsuffix cgo -o mini_webserv .

# Container for application run
FROM alpine:latest

RUN apk --no-cache add ca-certificates curl

# Adding POSTGRE
RUN apk --no-cache add postgresql
RUN (addgroup -S postgres && adduser -S postgres -G postgres || true)
RUN mkdir -p /var/lib/postgresql/data
RUN mkdir -p /run/postgresql/
RUN chown -R postgres:postgres /run/postgresql/
RUN chmod -R 777 /var/lib/postgresql/data
RUN chown -R postgres:postgres /var/lib/postgresql/data
RUN su - postgres -c "initdb /var/lib/postgresql/data"
RUN echo "host all  all    0.0.0.0/0  md5" >> /var/lib/postgresql/data/pg_hba.conf

RUN mkdir /app

WORKDIR /app/

COPY facts.sql .
COPY start.sh .
RUN chmod -R 777 start.sh

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /ozon/mini_webserv .

# Expose port for connection
EXPOSE 5432
EXPOSE 8888

# Run application
USER postgres
CMD ["./start.sh"]
