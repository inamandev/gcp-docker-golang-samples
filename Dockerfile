FROM golang:1.16-alpine

WORKDIR /app
COPY go.mod .
COPY src .
RUN go get
RUN go build -o hello-world main.go
EXPOSE 8081
CMD ["/app/hello-world"]