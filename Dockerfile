FROM golang:1.24-alpine as builder

RUN apk add --no-cache git make
WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Fixed: Added missing = sign and space in GOARCH
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o gothere .

FROM alpine:latest
RUN apk add --no-cache ca-certificates bash
COPY --from=builder /build/gothere /usr/local/bin/gothere

EXPOSE 443
ENTRYPOINT [ "/usr/local/bin/gothere", "relay", "-p", "443" ]