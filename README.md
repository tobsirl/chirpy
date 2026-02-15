# Chirpy

Chirpy is a simple Go web server that serves static files and provides basic metrics for homepage requests.

## Features

- Serves static files from the `/app/` and `/assets/` routes
- Tracks homepage requests and exposes a `/metrics` endpoint
- Provides a `/reset` endpoint to reset the request counter
- Health check endpoint at `/healthz`

## Usage

### Build and Run

You can use the provided Makefile:

```
make run
```

This will build the project and run the output binary from `./bin/out`.

### Endpoints

- `/app/` - Serves static files
- `/assets/` - Serves assets (e.g., images)
- `/metrics` - Returns the number of homepage requests since server start
- `/reset` - Resets the request counter to zero
- `/healthz` - Health check endpoint

## Requirements

- Go 1.19 or newer

## Project Structure

- `main.go` - Main server code
- `assets/` - Static assets
- `index.html` - Homepage
- `makefile` - Build and run commands

## License

MIT
