# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /app/

RUN mkdir ./logs

# Copy everything in the current directory into a Docker image
COPY . /app/

# The go mod download command downloads the named modules into the module cache.
RUN go mod download 

WORKDIR /app/cmd

# Compile application
RUN CGO_ENABLED=0 GOOS=linux go build -o /calculator-api .

EXPOSE 8080

# Run
CMD ["/calculator-api"]