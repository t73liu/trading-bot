# Build trader
FROM golang:1.14-alpine AS trader-builder

WORKDIR ${GOPATH}/src/trading-bot/trader/

COPY ../../trader .

RUN go mod download
RUN go mod verify

# Creates a static build: https://golang.org/cmd/link/
# Deprecate after: https://github.com/golang/go/issues/26492
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
  -ldflags='-s -w -extldflags "-static"' -a \
  -o /trader .

# Build dash
FROM node:lts-alpine AS dash-builder

COPY ../../dash .

RUN yarn install
RUN yarn run build

# Build image
FROM gcr.io/distroless/static-debian10:nonroot

COPY --from=dash-builder build /static
COPY --from=trader-builder /trader /

ENTRYPOINT ["/trader"]
