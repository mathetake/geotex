package geotex_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/mathetake/geotex"
	"github.com/mmcloughlin/geohash"
	"gotest.tools/assert"
)

func TestNewGeotex(t *testing.T) {
	for i, c := range []struct {
		acc     uint
		wantErr bool
	}{
		{
			acc:     4,
			wantErr: false,
		},
		{
			acc:     11,
			wantErr: false,
		},
		{
			acc:     0,
			wantErr: true,
		},
		{
			acc:     12,
			wantErr: true,
		},
	} {

		c := c
		t.Run(fmt.Sprintf("%d-th case", i), func(t *testing.T) {
			_, err := geotex.NewGeotex(c.acc)
			if c.wantErr {
				assert.Equal(t, true, err != nil)
			}
		})
	}
}

func TestVerifyQuarterLength(t *testing.T) {
	accToLen := geotex.GetExportedAccuracyToLength()

	for i, c := range []struct {
		count     float64
		threshold float64
		accuracy  uint
	}{
		{
			count:     1e5,
			threshold: 1e-10,
			accuracy:  2,
		},
		{
			count:     1e5,
			threshold: 1e-10,
			accuracy:  3,
		},
		{
			count:     1e5,
			threshold: 1e-10,
			accuracy:  4,
		},
		{
			count:     1e5,
			threshold: 1e-10,
			accuracy:  5,
		},
		{
			count:     1e5,
			threshold: 1e-10,
			accuracy:  6,
		},
		{
			count:     1e5,
			threshold: 1e-10,
			accuracy:  7,
		},
		{
			count:     1e5,
			threshold: 1e-10,
			accuracy:  8,
		},
		{
			count:     1e5,
			threshold: 1e-10,
			accuracy:  9,
		},
		{
			count:     1e5,
			threshold: 1e-10,
			accuracy:  10,
		},
		{
			count:     1e5,
			threshold: 1e-10,
			accuracy:  11,
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d-th case", i), func(t *testing.T) {
			ql, ok := accToLen[c.accuracy]
			if !ok {
				t.Fatal("invalid accuracy")
				return
			}

			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			var diffLat, diffLng, accLat, accLng float64
			for i := 0; i < int(c.count); i++ {
				lat := -90 + r.Float64()*180
				lng := -180 + r.Float64()*360
				h := geohash.EncodeWithPrecision(lat, lng, c.accuracy)
				box := geohash.BoundingBox(h)
				accLat += (box.MaxLat - box.MinLat) / 4
				accLng += (box.MaxLng - box.MinLng) / 4
				diffLat += math.Abs(ql.Lat - (box.MaxLat-box.MinLat)/4)
				diffLng += math.Abs(ql.Lng - (box.MaxLng-box.MinLng)/4)
			}

			assert.Equal(t, true, math.Abs(diffLat/c.count) < c.threshold)
			assert.Equal(t, true, math.Abs(diffLng/c.count) < c.threshold)
		})
	}
}

func TestGeotex_GetVertex(t *testing.T) {
	for i, c := range []struct {
		count int
		acc   uint
	}{
		{
			count: 1e3,
			acc:   2,
		},
		{
			count: 1e3,
			acc:   3,
		},
		{
			count: 1e3,
			acc:   4,
		},
		{
			count: 1e3,
			acc:   5,
		},
		{
			count: 1e3,
			acc:   6,
		},
		{
			count: 1e3,
			acc:   7,
		},
		{
			count: 1e3,
			acc:   8,
		},
		{
			count: 1e3,
			acc:   9,
		},
		{
			count: 1e3,
			acc:   10,
		},
		{
			count: 1e3,
			acc:   11,
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d-th case", i), func(t *testing.T) {
			g, err := geotex.NewGeotex(c.acc)
			if err != nil {
				t.Fatal(err)
				return
			}

			ql := geotex.GetQuarterLatLngFromGeotex(g)

			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			lat := -90 + r.Float64()*180
			lng := -180 + r.Float64()*360
			vLat, vLng := g.GetVertex(lat, lng)
			for j := 0; j < c.count; j++ {

				// check whether near points' vertex are same
				plat := r.Float64() * ql.Lat * 2
				plng := r.Float64() * ql.Lng * 2
				actualLat, actualLng := g.GetVertex(vLat+plat, vLng+plng)

				assert.Equal(t, vLat, actualLat)
				assert.Equal(t, vLng, actualLng)

				actualLat, actualLng = g.GetVertex(vLat+plat, vLng-plng)
				assert.Equal(t, vLat, actualLat)
				assert.Equal(t, vLng, actualLng)

				actualLat, actualLng = g.GetVertex(vLat-plat, vLng+plng)
				assert.Equal(t, vLat, actualLat)
				assert.Equal(t, vLng, actualLng)

				actualLat, actualLng = g.GetVertex(vLat-plat, vLng-plng)
				assert.Equal(t, vLat, actualLat)
				assert.Equal(t, vLng, actualLng)

				// check whether distant points' vertex are different
				plat = ql.Lat*2 + r.Float64()*ql.Lat*2
				plng = ql.Lng*2 + r.Float64()*ql.Lng*2
				actualLat, actualLng = g.GetVertex(vLat+plat, vLng+plng)

				assert.Equal(t, true, vLat != actualLat)
				assert.Equal(t, true, vLng != actualLng)

				actualLat, actualLng = g.GetVertex(vLat+plat, vLng-plng)
				assert.Equal(t, true, vLat != actualLat)
				assert.Equal(t, true, vLng != actualLng)

				actualLat, actualLng = g.GetVertex(vLat-plat, vLng+plng)
				assert.Equal(t, true, vLat != actualLat)
				assert.Equal(t, true, vLng != actualLng)

				actualLat, actualLng = g.GetVertex(vLat-plat, vLng-plng)
				assert.Equal(t, true, vLat != actualLat)
				assert.Equal(t, true, vLng != actualLng)
			}
		})
	}
}
