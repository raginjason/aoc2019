FROM golang:1.13-alpine as builder

# Install SSL ca certificates.
# Ca-certificates is required to call HTTPS endpoints.
# musl-dev is needed for gcc
RUN apk update && apk add --no-cache ca-certificates gcc musl-dev

WORKDIR /build

# Get dependencies setup first
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy rest of application in place and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main github.com/raginjason/aoc2019
RUN go test -v ./...

# Bare minimum container
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /app
COPY --from=builder /build/main .

CMD [ "/app/main" ]
