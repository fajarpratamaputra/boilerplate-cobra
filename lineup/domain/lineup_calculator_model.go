package domain

type (
	Content struct {
		ContentId int
		Title     string
	}

	Interaction struct {
		ContentId int
		UserId    int
		Action    string
		Service   string
	}
)
