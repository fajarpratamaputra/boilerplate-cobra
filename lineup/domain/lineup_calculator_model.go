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
)
