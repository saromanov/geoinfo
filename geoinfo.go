package geoinfo

import (
	"github.com/kellydunn/golang-geo"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
	"net/http"
	"time"
)

type Geoinfo struct {
	//Path to GetLite2-City.mmdb
	Path string
}

//IPDistance provides distance between two location by ip addresses
func (gi *Geoinfo) IPDistance(ip1, ip2 string) float64 {
	item1 := gi.getCity(ip1)
	item2 := gi.getCity(ip2)
	long1 := item1.Location.Longitude
	lat1 := item1.Location.Latitude
	long2 := item2.Location.Longitude
	lat2 := item2.Location.Latitude
	if long1 == long2 && lat1 == lat2 {
		return 0
	}

	p := geo.NewPoint(long1, lat1)
	p2 := geo.NewPoint(long2, lat2)

	return p.GreatCircleDistance(p2)
}

//GetCountryByIP return country in en. format
func (gi *Geoinfo) GetCountryByIP(ip string) string {
	item := gi.getCity(ip)
	return item.Country.Names["en"]
}

func (gi *Geoinfo) GetPointByIP(ip string) (*geo.Point, error) {
	item := gi.getCity(ip)
	geocoder := geo.GoogleGeocoder{&http.Client{Timeout: time.Duration(5 * time.Second)}}
	return geocoder.Geocode(item.Country.Names["en"])
}

func (gi *Geoinfo) getCity(ip string) *geoip2.City {
	db, err := geoip2.Open(gi.Path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	validip := net.ParseIP(ip)
	record, err := db.City(validip)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(record.Location.Latitude, record.Location.Longitude)
	return record
}
