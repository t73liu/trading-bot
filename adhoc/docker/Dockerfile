# Build trader
FROM golang:1.18-alpine AS trader-builder

WORKDIR ${TRADING_BOT_REPO}/trader/

COPY . .

RUN go mod download

# Creates a static build: https://golang.org/cmd/link/
# Deprecate after: https://github.com/golang/go/issues/26492
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
  -ldflags='-s -w -extldflags "-static"' -a \
  -o /go/bin/trader

# Build dash
FROM node:lts-alpine AS dash-builder

WORKDIR ${TRADING_BOT_REPO}/dash/

COPY . .

RUN npm install -g pnpm
RUN pnpm install
RUN pnpm run build

# Build image
FROM gcr.io/distroless/static-debian11

COPY --from=dash-builder dist /static
COPY --from=trader-builder /go/bin/trader /

ENTRYPOINT ["/trader"]
