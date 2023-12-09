FROM golang:alpine AS builder

WORKDIR /build

COPY ./algopack .

RUN go build -o task ./cmd/algopack/main

FROM alpine

WORKDIR /build

COPY --from=builder /build/task /build/task

CMD ["./task", "-t", "60"]