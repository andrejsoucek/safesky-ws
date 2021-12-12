package aircraft

import "github.com/andrejsoucek/safesky-ws/geography"

type Aircraft struct {
	Id               string           `json:"id"`
	Source           string           `json:"source"`
	TransponderType  string           `json:"transponderType"`
	BeaconType       string           `json:"beaconType"`
	LastUpdate       float64          `json:"lastUpdate"`
	LatLng           geography.LatLng `json:"latLng"`
	Altitude         float64          `json:"altitude"`
	VerticalRate     float64          `json:"verticalRate"`
	Accuracy         float64          `json:"accuracy"`
	AltitudeAccuracy float64          `json:"altitudeAccuracy"`
	Course           float64          `json:"course"`
	GroundSpeed      float64          `json:"groundSpeed"`
	Status           string           `json:"status"`
	TurnRate         float64          `json:"turnRate"`
	CallSign         string           `json:"callSign"`
}

func CreateFromResponse(resp []interface{}) Aircraft {
	return Aircraft{
		Id:               toString(resp[0]),
		Source:           toString(resp[1]),
		TransponderType:  toString(resp[2]),
		BeaconType:       toString(resp[3]),
		LastUpdate:       toFloat(resp[4]),
		LatLng:           geography.LatLng{Lat: toFloat(resp[5]), Lng: toFloat(resp[6])},
		Altitude:         toFloat(resp[7]),
		VerticalRate:     toFloat(resp[8]),
		Accuracy:         toFloat(resp[9]),
		AltitudeAccuracy: toFloat(resp[10]),
		Course:           toFloat(resp[11]),
		GroundSpeed:      toFloat(resp[12]),
		Status:           toString(resp[13]),
		TurnRate:         toFloat(resp[14]),
		CallSign:         toString(resp[15]),
	}
}

func toString(v interface{}) string {
	if v == nil {
		return ""
	}
	return v.(string)
}

func toFloat(v interface{}) float64 {
	if v == nil {
		return 0
	}
	return v.(float64)
}
