# Control Plane

Go-based management API for Nginx AI Gateway.

## Features

- API key management
- Model configuration
- Analytics dashboard
- Nginx Plus API integration

## Development on macOS

```bash
# Install dependencies
go mod download

# Run locally
go run cmd/server/main.go

# Build
go build -o control-plane cmd/server/main.go

# Or use air for hot reload
brew install air
air
```

## API

Runs on http://localhost:8000

Endpoints:
- GET /health
- GET /api/v1/models
- POST /api/v1/api-keys
- GET /api/v1/analytics
