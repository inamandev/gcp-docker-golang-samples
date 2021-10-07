FROM golang:1.16-alpine

WORKDIR /app
COPY src .
RUN go get
RUN go build -o hello-world main.go
CMD ["/app/hello-world"]