package redis

type LineupPayload struct {
	ContentId   interface{} `json:"content_id"`
	Service     string      `json:"service"`
	ContentType string      `json:"content_type"`
}
