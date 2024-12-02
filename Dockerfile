FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o crm.goshop.com ./cmd/server

FROM scratch

COPY ./config /config

COPY --from=builder /build/crm.goshop.com /

ENTRYPOINT ["/crm.goshop.com", "/config/local.yaml"]