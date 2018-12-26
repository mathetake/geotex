package geotex

type ExportedQuarterLength struct {
	Lat, Lng float64
}

func GetExportedAccuracyToLength() map[uint]*ExportedQuarterLength {
	ret := make(map[uint]*ExportedQuarterLength, len(accuracyToLength))
	for k, v := range accuracyToLength {
		ret[k] = &ExportedQuarterLength{
			Lat: v.lat,
			Lng: v.lng,
		}
	}
	return ret
}

func GetQuarterLatLngFromGeotex(g *Geotex) *ExportedQuarterLength {
	return &ExportedQuarterLength{
		Lat: g.quarterLength.lat,
		Lng: g.quarterLength.lng,
	}
}
