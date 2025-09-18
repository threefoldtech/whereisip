package main

import (
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
)

type ipresponse struct {
	Address     string  `json:"address"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
	Subdivision string  `json:"subdivision"`
	Continent   string  `json:"continent"`
	CityName    string  `json:"city_name"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Source      string  `json:"source"`
}

func main() {
	db, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	err = r.SetTrustedProxies(nil) // nil trusts all proxies
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	r.GET("/", func(c *gin.Context) {
		ip := "0.0.0.0"

		if c.Request.Header.Get("X-Real-IP") != "" {
			ip = c.Request.Header.Get("X-Real-IP")

		} else {
			ip, _, err = net.SplitHostPort(c.Request.RemoteAddr)

			if err != nil {
				log.Printf("debug: Getting req.RemoteAddr %v", err)
				c.JSON(500, gin.H{"error": "could not determine ip address"})
			}
		}

		ipobj := net.ParseIP(ip)
		record, err := db.City(ipobj)

		if err != nil {
			log.Printf("%v", err)
			c.JSON(500, gin.H{"error": "db query failed"})
		}

		response := ipresponse{
			ip,
			record.Country.IsoCode,
			"Unknown",
			"Unknown",
			"Unknown",
			"Unknown",
			record.Location.Latitude,
			record.Location.Longitude,
			"maxmind geolite2", // for credits
		}

		if val, ok := record.Country.Names["en"]; ok {
			response.CountryName = val
		}

		if val, ok := record.Continent.Names["en"]; ok {
			response.Continent = val
		}

		if val, ok := record.City.Names["en"]; ok {
			response.CityName = val
		}

		if len(record.Subdivisions) > 0 {
			response.Subdivision = record.Subdivisions[0].Names["en"]
		}

		c.JSON(200, response)
	})

	log.Print("Starting web server")
	r.Run(":80")
}
