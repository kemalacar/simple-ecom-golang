# Latest golang image on apline linux
FROM golang:1.17 as builder

# Env variables
ENV GOOS linux
ENV CGO_ENABLED 0

# Work directory
WORKDIR /golang-my-market-api

# Installing dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copying all the files
COPY . .

# Building the application
RUN go build -o golang-my-market-api

# Fetching the latest nginx image
FROM alpine:3.16 as production

# Certificates
RUN apk add --no-cache ca-certificates

# Copying built assets from builder
COPY --from=builder golang-my-market-api .

# Starting our application
CMD ./golang-my-market-api

# Exposing server port
EXPOSE 5000