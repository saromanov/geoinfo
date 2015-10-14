# geoinfo

Based on MaxMind (GeoLite2)[http://dev.maxmind.com/geoip/geoip2/geolite2/]

### Installation
``` go get github.com/saromanov/geoinfo ```


### API

### Example
```go
package main

import (
  "github.com/saromanov/geoinfo"
  "fmt"
)

func main() {
	gi := geoinfo.Geoinfo{Path: "../GeoLite2-City.mmdb"}
	fmt.Println(gi.IPDistance("87.240.131.118", "213.180.204.3"))
}
```

### License
MIT

