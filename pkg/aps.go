package pkg

type Payload struct {
	Alert string 	`json:"alert"`
	Sound string 	`json:"sound"`
	LinkUrl string	`json:"link_url"`
}

type APS struct {
	APS Payload		`json:"aps"`
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
