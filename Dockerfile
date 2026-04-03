# syntax=docker/dockerfile:1

# build stage
FROM golang:1.25.8 AS builder

WORKDIR /permit-proxy

COPY permit-proxy/ .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o go-sj-permits

# run stage
FROM gcr.io/distroless/static

COPY --from=builder /permit-proxy/go-sj-permits /permit-proxy

CMD ["/permit-proxy"]
