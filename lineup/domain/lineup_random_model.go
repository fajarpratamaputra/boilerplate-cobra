package domain

type ContentRandom struct {
	ContentId   int         `json:"content_id"`
	Service     interface{} `json:"service"`
	ContentType interface{} `json:"content_type"`
}

type ContentString struct {
	ContentId   int         `json:"content_id"`
	Service     interface{} `json:"service"`
	ContentType interface{} `json:"content_type"`
}
