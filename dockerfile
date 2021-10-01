FROM golang:1.16.6-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
ENV GO111MODULE=on
RUN go mod download
RUN go build -o main .
CMD ["./main"]

