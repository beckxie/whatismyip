# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.2.0] - 2026-01-04

### Added
- **API endpoint** `/api/ip` for programmatic access
  - Plain text format: `curl /api/ip`
  - JSON format with detailed info: `curl "/api/ip?format=json"`
- CORS headers support for cross-origin requests
- IP version detection (IPv4/IPv6)
- GitHub Actions CI workflow

### Changed
- Restructured to standard Go project layout (`cmd/`, `internal/`)
- Improved Makefile with more build targets
- Updated README with API documentation
- Fixed docker-compose.yaml configuration errors

### Technical
- Added `internal/api/` package for API handlers
- Added `internal/web/` package for web handlers
- Separated response types into `internal/api/response.go`

## [1.1.0] - 2026-01-04

### Changed
- Migrated to Go 1.23
- Updated logging to use `slog`
- Updated Dockerfile to use distroless image
- Improved IP detection logic with proper validation

## [1.0.0] - 2020-xx-xx

### Added
- Initial release
- Web interface showing public IP address
- HTTP headers display
- Nginx/Caddy reverse proxy examples
