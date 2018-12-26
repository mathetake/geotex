# geotex [![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

Given a latlng, __geotex__ can be used to get the vertex's of its geohash's rectangle  and get geohashes whose rectangle contains the vertex.

# usage

```golang

package main

import (
	"fmt"

	"github.com/mathetake/geotex"
)

func main() {
	var accuracy uint = 6
	g, _ := geotex.NewGeotex(accuracy)

	vLat, vLng := g.GetVertex(0.1, 0.1)

	fmt.Printf("the nearest vertex's (lat, lng) = (%f, %f)", vLat, vLng)
}

```


# License
MIT