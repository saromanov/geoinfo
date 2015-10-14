package geoinfo

import (
   "testing"
)

func TestIPDistance(t *testing.T) {
	info := Geoinfo{Path: "../GeoLite2-City.mmdb"}
	result1 := info.IPDistance("87.240.131.118", "81.2.69.142")
	expected1 := 4318.067046451507
	if result1 != expected1 {
		t.Errorf("Expected %d, found - %d", expected1, result1)
	}

	result2 := info.IPDistance("87.240.131.118", "213.180.204.3")
	expected2 := 0.0
	if result2 != expected2 {
		t.Errorf("Expected %d, found - %d", expected2, result2)
	}

}

func TestGetCountryByIP(t *testing.T) {
	info := Geoinfo{Path: "../GeoLite2-City.mmdb"}
	result1 := info.GetCountryByIP("87.240.131.118")
	expected := "Russia"
	if result1 != expected {
		t.Errorf("Expected %s is not equal to %s", expected, result1)
	}
}