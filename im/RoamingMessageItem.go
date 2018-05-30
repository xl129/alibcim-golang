package im

type RoamingMessageItem struct {
	Type string `json:"type"`
	Value string `json:"value"`
}

type RoamingMessage struct {
	ContentList []RoamingMessage `json:"content_list"`
	Time int64 `json:"time"`
	Direction int `json:"direction"`
	Uuid int64 `json:"uuid"`
	Type int `json:"type"`
}