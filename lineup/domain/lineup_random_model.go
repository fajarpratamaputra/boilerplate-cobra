package domain

type ContentRandom struct {
	ContentId   interface{} `json:"content_id"`
	Service     interface{} `json:"service"`
	ContentType interface{} `json:"content_type"`
}

type ContentString struct {
	ContentId   interface{} `json:"content_id"`
	Service     interface{} `json:"service"`
	ContentType interface{} `json:"content_type"`
}
