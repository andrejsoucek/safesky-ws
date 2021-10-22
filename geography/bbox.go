package geography

type BoundingBox struct {
	SouthWest LatLon `json:"sw"`
	NorthEast LatLon `json:"ne"`
}

func IsInBounds(bb BoundingBox, p LatLon) bool {
	ne := bb.NorthEast
	sw := bb.SouthWest
	isLonInRange := false

	if ne.Lon < sw.Lon {
		isLonInRange = p.Lon >= sw.Lon || p.Lon <= ne.Lon
	} else {
		isLonInRange = p.Lon >= sw.Lon && p.Lon <= ne.Lon
	}

	return isLonInRange && p.Lat >= sw.Lat && p.Lat <= ne.Lat
}
