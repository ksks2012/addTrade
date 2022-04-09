# introduce

[x] stream agg trade from websocket API
[x] websocket server for client
[x] gin API to get newest data

## structure

- client: redis & websocket client
- etc: setting file
- docs: document
- global: global variables
- internal (internal module):
  <!-- TODO: -->
  - dao: data access object
  - middleware
  - model: database model control
  - routers: api routes
  - service: process business logic

- pkg: package
- storage: temp file
- server: websocket server
- scripts: build, install, analysis scripts
- third_party: third_party tools

# Subscript ws

<https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-streams_CN.md>

# Redis

default setting

# build

## main service

```bash
go build ./cmd/add_trade
```

## benchmark client for websocket server

```bash
go build ./cmd/benchmark
```

# websocket server

```
/stream
  token=<token>
```

# gin API

```
/AggTrade
 content=btcusdt
```
