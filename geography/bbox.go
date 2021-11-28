package geography

import "encoding/json"

type BoundingBox struct {
	SouthWest LatLng `json:"sw"`
	NorthEast LatLng `json:"ne"`
}

func (bb BoundingBox) Contains(p LatLng) bool {
	ne := bb.NorthEast
	sw := bb.SouthWest
	isLonInRange := false

	if ne.Lng < sw.Lng {
		isLonInRange = p.Lng >= sw.Lng || p.Lng <= ne.Lng
	} else {
		isLonInRange = p.Lng >= sw.Lng && p.Lng <= ne.Lng
	}

	return isLonInRange && p.Lat >= sw.Lat && p.Lat <= ne.Lat
}

func CreateBoundingBoxFromJson(data string) BoundingBox {
	bb := BoundingBox{}
	json.Unmarshal([]byte(data), &bb)

	return bb
}
