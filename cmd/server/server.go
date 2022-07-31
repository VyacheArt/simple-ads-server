package main

import (
	"github.com/VyacheslavGoryunov/simple-ads-server/internal/ads"
	"github.com/VyacheslavGoryunov/simple-ads-server/internal/stats"
	"github.com/VyacheslavGoryunov/simple-ads-server/internal/stats/mysql"
	"github.com/oschwald/geoip2-golang"
	"log"
	"time"
)

func main() {
	reader, err := geoip2.Open("GeoLite2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}

	mw, err := mysql.NewMySqlWriter("127.0.0.1", 13306, "rotator", "statistics",
		"root", "qwerty123")
	if err != nil {
		log.Fatal(err)
	}

	statsManager := stats.NewManager(mw, 10*time.Second)
	statsManager.Start()

	s := ads.NewServer(reader, statsManager)

	if err := s.Listen(); err != nil {
		log.Fatal(err)
	}
}
