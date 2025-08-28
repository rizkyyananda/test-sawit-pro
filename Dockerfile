FROM golang:1.23.10 AS builder
WORKDIR /src

# depedency-injection sini go.mod berada depedency-injection konteks ini (test_sawit_pro)
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
RUN go build -o /out/app .

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /out/app /app/app
EXPOSE 9000
ENTRYPOINT ["/app/app"]