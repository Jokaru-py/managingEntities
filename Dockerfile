FROM golang:1.16.0-alpine3.13

# RUN go version
ENV GOPATH=/

COPY ./ ./

# build go app
# RUN go mod download
RUN go build -o main ./cmd/managing-entities-server/main.go

# WORKDIR /
CMD ["./main"]