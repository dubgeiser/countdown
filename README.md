# Countdown
An implementation (very much WIP and POC) of the game "Countdown", using plain Go, HTML, CSS and JavaScript.
There isn't yet a build system, so it's just `go run ...` for the time being.

## Build/run
Clone this repository.
Run the server via `go run cmd/server/main.go` or `make` this will start the server on the default port, which is `8080`.

To run the server on a different port, you can pass the `--port` option, see `go run cmd/server/main.go --help`.
For instance: `go run cmd/server/main.go --port=6969` will run the server on port `6969`.

When the server has started, point your browser to `localhost:8080` (or whatever port has been passed with the `--port` option).
