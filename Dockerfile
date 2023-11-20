#Simple Docker File

FROM golang:1.13
EXPOSE 3000
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get -u github.com/lib/pq
RUN go build -o main .
RUN go build -o mainworker worker/main.go
ENTRYPOINT ["/app/main"]
