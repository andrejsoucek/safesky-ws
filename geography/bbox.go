package geography

type BoundingBox struct {
	SouthWest LatLng
	NorthEast LatLng
}

func IsInBounds(bb BoundingBox, p LatLng) bool {
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
