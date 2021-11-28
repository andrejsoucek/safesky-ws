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
		Id:               resp[0].(string),
		Source:           resp[1].(string),
		TransponderType:  resp[2].(string),
		BeaconType:       resp[3].(string),
		LastUpdate:       resp[4].(float64),
		LatLng:           geography.LatLng{Lat: resp[5].(float64), Lng: resp[6].(float64)},
		Altitude:         resp[7].(float64),
		VerticalRate:     resp[8].(float64),
		Accuracy:         resp[9].(float64),
		AltitudeAccuracy: resp[10].(float64),
		Course:           resp[11].(float64),
		GroundSpeed:      resp[12].(float64),
		Status:           resp[13].(string),
		TurnRate:         resp[14].(float64),
		CallSign:         resp[15].(string),
	}
}
