# Grid GeoIP

Grid GeoIP is a lightweight Go webservice that provides IP geolocation and lookup. It uses the MaxMind GeoLite2 database to determine geographical location and metadata associated with IP addresses.

## What this is

Grid GeoIP is a small HTTP service that answers geolocation queries for IP addresses. It supports both explicit IP queries via request parameters and implicit lookups based on the remote address of the incoming request. The service supports both IPv4 and IPv6 addresses and returns structured JSON containing country, city, continent, subdivision, and coordinate data.

## What this repository contains

- `main.go` — Go source for the HTTP service using the Gin web framework
- `go.mod` / `go.sum` — Go module definition and dependency lock
- `update.sh` — Script for updating the GeoLite2 database

## Role in the stack

Grid GeoIP functions as a network diagnostics and geolocation utility. It can be deployed as a standalone microservice or integrated into larger systems that need to determine the geographical origin of traffic. It is commonly used for node geography mapping, traffic analysis, and regional service routing.

## Relation to ThreeFold

This technology is used within the ThreeFold ecosystem and was first deployed on the ThreeFold Grid. The component itself is designed as reusable infrastructure technology and should be understood by its technical function first, independent of any specific deployment.

## Ownership

This repository is owned and maintained by TF-Tech NV, a Belgian company responsible for the development and maintenance of this technology.

## Build

```bash
go build
```

## Installation

1. Obtain a copy of the MaxMind GeoLite2-City database (`GeoLite2-City.mmdb`).
2. Place the database file in the same directory as the `whereisip` binary.
3. Run the service: `./whereisip`

The service listens on port `80` by default. There are no command-line arguments.

## Usage

Send a GET request to `/`. You can optionally provide an `ip` query parameter.

```bash
curl "http://localhost/?ip=8.8.8.8"
```

Example response:

```json
{
  "address": "8.8.8.8",
  "country_code": "US",
  "country_name": "United States",
  "subdivision": "California",
  "continent": "North America",
  "city_name": "Mountain View",
  "latitude": 37.386,
  "longitude": -122.0838,
  "source": "maxmind geolite2"
}
```

## Dependencies

- [MaxMind GeoLite2](https://www.maxmind.com/en/geoip2-city)
- [Gin web framework](https://github.com/gin-gonic/gin)
- [geoip2-golang](https://github.com/oschwald/geoip2-golang)

## License

This project is licensed under the Apache License 2.0 — see the [LICENSE](LICENSE) file for details.
