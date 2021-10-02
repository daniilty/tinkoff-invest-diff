FROM golang:1.17.0-buster AS builder

WORKDIR "/app"

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY cmd cmd
COPY internal internal
COPY Makefile .

RUN make build

FROM debian:buster

USER nobody

COPY --from=builder /app/diff /app/diff

ENTRYPOINT ["/app/diff"]
