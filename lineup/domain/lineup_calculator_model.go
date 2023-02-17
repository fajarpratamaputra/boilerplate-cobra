package domain

type (
	Content struct {
		ContentId   int
		Service     string
		ContentType string
		Score       float64
	}

	Lineup map[interface{}]*Content
)
