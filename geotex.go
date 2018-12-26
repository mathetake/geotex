package geotex

import (
	"github.com/mmcloughlin/geohash"
	"github.com/pkg/errors"
)

var (
	accuracyToLength = map[uint]*quarterLength{
		2:  {lat: 1.40625, lng: 2.8125},
		3:  {lat: 0.3515625, lng: 0.3515625},
		4:  {lat: 0.0439453125, lng: 0.087890625},
		5:  {lat: 0.010986328125, lng: 0.010986328125},
		6:  {lat: 0.001373291015625, lng: 0.00274658203125},
		7:  {lat: 0.00034332275390625, lng: 0.00034332275390625},
		8:  {lat: 4.291534423828125e-05, lng: 8.58306884765625e-05},
		9:  {lat: 1.0728836059570312e-05, lng: 1.0728836059570312e-05},
		10: {lat: 1.341104507446289e-06, lng: 2.682209014892578e-06},
		11: {lat: 3.3527612686157227e-07, lng: 3.3527612686157227e-07},
	}
)

type Geotex struct {
	quarterLength *quarterLength
	accuracy      uint
}

type quarterLength struct {
	lat, lng float64
}

func NewGeotex(acc uint) (*Geotex, error) {
	if ql, ok := accuracyToLength[acc]; ok {
		return &Geotex{quarterLength: ql, accuracy: acc}, nil
	}

	validAcc := make([]uint, 0, len(accuracyToLength))
	for acc := range accuracyToLength {
		validAcc = append(validAcc, acc)
	}
	return nil, errors.Errorf("invalid accuracy not in :%v", validAcc)
}

func (g *Geotex) GetVertex(lat, lng float64) (float64, float64) {
	gh := geohash.EncodeWithPrecision(lat, lng, g.accuracy)
	box := geohash.BoundingBox(gh)

	var retLng = box.MaxLng
	if (lng - box.MinLng) < (box.MaxLng - lng) {
		retLng = box.MinLng
	}

	var retLat = box.MaxLat
	if (lat - box.MinLat) < (box.MaxLat - lat) {
		retLat = box.MinLat
	}
	return retLat, retLng
}

func (g *Geotex) GetNearestRectangleInHash(rLat, rLng float64) []string {
	return []string{
		geohash.EncodeWithPrecision(rLat+g.quarterLength.lat, rLng+g.quarterLength.lng, g.accuracy),
		geohash.EncodeWithPrecision(rLat-g.quarterLength.lat, rLng+g.quarterLength.lng, g.accuracy),
		geohash.EncodeWithPrecision(rLat+g.quarterLength.lat, rLng-g.quarterLength.lng, g.accuracy),
		geohash.EncodeWithPrecision(rLat-g.quarterLength.lat, rLng-g.quarterLength.lng, g.accuracy),
	}
}
