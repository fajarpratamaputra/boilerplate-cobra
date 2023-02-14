package mongo

type (
	InteractionModel struct {
		ContentID int    `bson:"contentid"`
		Action    string `bson:"action"`
	}
)
