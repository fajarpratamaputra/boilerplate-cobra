package main

type (
	content struct {
		contentId int
		title     string
	}

	interaction struct {
		contentId int
		userId    int
		action    string
		service   string
	}
)
