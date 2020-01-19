package apnian

import "github.com/electronicpanopticon/gobrick"

type Payload struct {
	Alert   string `json:"alert"`
	Sound   string `json:"sound"`
	LinkUrl string `json:"link_url"`
}

type APS struct {
	APS Payload `json:"aps"`
}

func GenerateAPS(alert string, sound string, linkUrl string) *APS {
	payload := Payload{
		Alert:   alert,
		Sound:   sound,
		LinkUrl: linkUrl,
	}
	return &APS{
		APS: payload,
	}
}

func (aps APS) ToJsonBytes() []byte {
	return gobrick.ToJSONBytes(aps)
}
