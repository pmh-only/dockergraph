FROM golang:alpine as builder
WORKDIR /build

RUN apk add --no-cache gcc
RUN apk add --no-cache musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=1 go build

FROM alpine
WORKDIR /app
USER 1000:1000

COPY --from=builder /build/dockergraph .
CMD ["./dockergraph"]
