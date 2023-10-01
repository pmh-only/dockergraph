FROM golang:alpine

RUN go build .

CMD ["./dockergraph"]
