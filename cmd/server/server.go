package main

import (
	"github.com/oschwald/geoip2-golang"
	"log"
	"mini-ads-server1/internal/ads"
)

func main() {
	reader, err := geoip2.Open("GeoLite2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}

	s := ads.NewServer(reader)

	if err := s.Listen(); err != nil {
		log.Fatal(err)
	}
}
