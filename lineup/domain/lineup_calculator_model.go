package domain

type (
	Content struct {
		ContentId int
		Title     string
	}

	LineupContent struct {
		ContentId   int
		Service     string
		ContentType string
		Score       float64
	}

	LineupMap map[interface{}]*LineupContent

	LineupPayload struct {
		ContentId   interface{} `json:"content_id"`
		Service     string      `json:"service"`
		ContentType string      `json:"content_type"`
	}
)
