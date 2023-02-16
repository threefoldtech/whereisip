# Where Is My IP

Small webservice using GeoLite2 database to give geographic information about IP address location

# Dependencies

This project depends on [MaxMind GeoLite2](https://www.maxmind.com/en/geoip2-city)

# Build

Simply build the project with Go:
```go
go build
```

# Installation

First you need to copy a version of `GroLite2-City` database from MaxMind.
Copy this file into the same directory than `whereisip` binary: `GeoLite2-City.mmdb`

Then run the webservice: `./whereisip`

There is no command line arguments supported for now.

# Usage

Only one GET request to `/` is supported and you'll get a JSON response of GeoLite database content based on your remote address.
Server listen on `:8085`.

Example response:
```json
{
  "address": "82.212.171.24.",
  "country_code": "BE",
  "country_name": "Belgium",
  "subdivision": "Wallonia",
  "continent": "Europe",
  "city_name": "Liège",
  "latitude": 50.6337,
  "longitude": 5.5675,
  "source": "maxmind geolite2"
}
```
