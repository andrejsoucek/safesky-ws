package geography

import "encoding/json"

type BoundingBox struct {
	SouthWest LatLon `json:"sw"`
	NorthEast LatLon `json:"ne"`
}

func (bb BoundingBox) Contains(p LatLon) bool {
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

func CreateBoundingBoxFromJson(data string) BoundingBox {
	bb := BoundingBox{}
	json.Unmarshal([]byte(data), &bb)

	return bb
}
